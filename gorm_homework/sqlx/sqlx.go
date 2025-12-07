// 题目1：使用SQL扩展库进行查询
// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

// 题目2：实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
package main

import (
	"fmt"
	"log"

	// 纯 Go SQLite 驱动
	"github.com/glebarez/sqlite"
	"github.com/jmoiron/sqlx"
)

type Employees struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

func createEmpTable(db *sqlx.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS employees (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        department TEXT,
        salary INTEGER
    )`
	// id 自增主键值
	// name 非空string
	// department string
	// salary int

	_, err := db.Exec(sql)
	if err != nil {
		log.Fatal("创建表失败:", err)
	}
	fmt.Println("表创建/检查完成")
}

func insertEmpData(db *sqlx.DB) {
	var count int
	db.Get(&count, "SELECT COUNT(*) FROM employees")

	if count == 0 {
		// 简单方式插入数据
		db.MustExec(`INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)`,
			"张三", "技术部", 15000)
		db.MustExec(`INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)`,
			"李四", "技术部", 18000)
		db.MustExec(`INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)`,
			"王五", "市场部", 12000)
		db.MustExec(`INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)`,
			"赵六", "人事部", 10000)

		fmt.Println("测试数据插入完成")
	} else {
		fmt.Println("表中已有数据")
	}
}

type Book struct {
	ID     int    `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	Price  int    `db:"price"`
}

func main() {
	// 使用纯 Go 的 SQLite 驱动
	db, err := sqlx.Open(sqlite.DriverName, "./sqlx.db")
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	//sqlx 必须显式调用 db.Close()
	defer db.Close()

	fmt.Println("数据库连接成功!")

	createEmpTable(db)
	insertEmpData(db)
	var system_employees []Employees
	// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	err = db.Select(&system_employees,
		"SELECT * FROM employees WHERE department = ?", "技术部")

	if err != nil {
		log.Fatal("查询失败:", err)
	}

	fmt.Printf("\n找到 %d 个技术部员工:\n", len(system_employees))
	for _, emp := range system_employees {
		fmt.Printf("  - ID:%d 姓名:%s 薪资:%d\n", emp.ID, emp.Name, emp.Salary)
	}

	var best_employee Employees
	// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体
	db.Get(&best_employee, "SELECT * FROM employees ORDER  BY salary DESC LIMIT 1")
	fmt.Printf("  薪资最高的员工是ID:%d 姓名:%s 薪资:%d\n", best_employee.ID, best_employee.Name, best_employee.Salary)

	// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

	createBookTable(db)
	insertBookData(db)

	// 定义一个 Book 结构体，包含与 books 表对应的字段。
	// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
	var books []Book

	db.Select(&books, "SELECT * FROM books WHERE price > ?", 50)
	for _, book := range books {
		fmt.Printf("book 表中 价格超过 50 元的 ID:%d title:%s author:%s price:%d\n", book.ID, book.Title, book.Author, book.Price)
	}
}
func createBookTable(db *sqlx.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS books (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        author TEXT,
        price INTEGER
    )`
	// id 自增主键值
	// title 非空string
	// author string
	// price int

	_, err := db.Exec(sql)
	if err != nil {
		log.Fatal("创建表失败:", err)
	}
	fmt.Println("表创建/检查完成")
}

func insertBookData(db *sqlx.DB) {
	var count int
	db.Get(&count, "SELECT COUNT(*) FROM books")

	if count == 0 {
		// 简单方式插入数据
		db.MustExec(`INSERT INTO books (title, author, price) VALUES (?, ?, ?)`,
			"哈利波特", "jk", 60)
		db.MustExec(`INSERT INTO books (title, author, price) VALUES (?, ?, ?)`,
			"哈利波特2", "jk", 16)
		db.MustExec(`INSERT INTO books (title, author, price) VALUES (?, ?, ?)`,
			"哈利波特3", "jk", 26)
		db.MustExec(`INSERT INTO books (title, author, price) VALUES (?, ?, ?)`,
			"哈利波特4", "jk", 66)

		fmt.Println("测试数据插入完成")
	} else {
		fmt.Println("表中已有数据")
	}
}
