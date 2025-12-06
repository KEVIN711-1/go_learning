// 假设有两个表:

// 	accounts 表（包含字段 id 主键， balance 账户余额）
// 	和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。

// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
// 在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
// 并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Accounts struct {
	Name    string
	ID      uint `gorm:"primaryKey;autoIncrement"` // 自动递增
	Balance int
}

type Transactions struct {
	ID              uint `gorm:"primaryKey;autoIncrement"` // 自动递增
	Amount          int
	From_account_id int
	To_account_id   int
}

// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
// 在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
// 并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
func transfer_account(db *gorm.DB, from_id int, to_id int, trans_amount int) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// A账户扣钱
		// 为什么在事务中，不直接操作tb_act,而是使用tx?
		// 因为，直接操作tb_act，成功会立即生效
		// 但是操作tx的，只要有一个tx操作失败，所有的操作会回滚，这也是事务的一个特性
		if err := tx.Model(&Accounts{}).
			Where("ID = ? AND Balance >= ?", from_id, trans_amount).
			Update("Balance", gorm.Expr("Balance - ?", trans_amount)).Error; err != nil {
			return err
		}

		// B账户加钱
		fmt.Printf(" working here 2\n")

		if err := tx.Model(&Accounts{}).
			Where("ID = ?", to_id).
			Update("Balance", gorm.Expr("Balance + ?", trans_amount)).Error; err != nil {
			return err
		}

		transactions := []Transactions{
			{
				Amount:          trans_amount,
				From_account_id: from_id,
				To_account_id:   to_id,
			},
		}

		// 并在 transactions 表中记录该笔转账信息。
		tx.Create(&transactions)
		tx.Find(&transactions)

		fmt.Printf("数据库中有 %d 笔交易记录\n", len(transactions))
		for _, s := range transactions {
			fmt.Printf("ID = %d, 金额: %d from_id: %d to_id:%d\n", s.ID, s.Amount, s.From_account_id, s.To_account_id)
		}
		//err 失败 只要失败，事务所有操作回滚，要么全部成功，要么全部失败，防止转账造成的钱财丢失
		//nil 成功
		return nil
	})

}

func main() {
	// 1. 连接数据库
	db_act, _ := gorm.Open(sqlite.Open("./account.db"), &gorm.Config{})
	db_act.Exec("DELETE FROM account")

	db_act.AutoMigrate(&Accounts{}, &Transactions{})

	accounts := []Accounts{
		{Name: "zkf", Balance: 1000},
		{Name: "lgy", Balance: 400},
	}

	db_act.Create(&accounts)

	fmt.Printf("有%d 账户\n", len(accounts))
	for _, s := range accounts {
		fmt.Printf("ID = %d, 用户名; %s 余额: %d) \n", s.ID, s.Name, s.Balance)
	}

	// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
	transfer_account(db_act, 1, 2, 100)

	// 4. 重新查询数据库获取最新数据
	fmt.Println("\n=== 转账后最新状态 ===")
	var latestAccounts []Accounts

	// 	数据库更新了，但内存变量没更新
	// Find() 会从数据库读取最新数据
	// 这样你看到的就是真实的数据库状态
	db_act.Find(&latestAccounts)
	fmt.Printf("数据库中有 %d 账户\n", len(latestAccounts))
	for _, s := range latestAccounts {
		fmt.Printf("ID = %d, 用户名: %s 余额: %d\n", s.ID, s.Name, s.Balance)
	}
}
