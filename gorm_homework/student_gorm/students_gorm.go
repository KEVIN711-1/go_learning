// 假设有一个名为 students 的表，
// 包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
// 要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Student struct {
	ID    uint `gorm:"primaryKey;autoIncrement"` // 自动递增
	Name  string
	Grade string
	Age   int
}

func main() {
	// 1. 连接数据库
	db, _ := gorm.Open(sqlite.Open("./students.db"), &gorm.Config{})
	//创建不存在的表
	// 添加不存在的字段
	// 创建索引
	// 不会删除已有数据或字段
	db.Exec("DELETE FROM students")

	db.AutoMigrate(&Student{})

	students := []Student{
		{Name: "zkf", Grade: "4年级", Age: 13},
		{Name: "lgy", Grade: "5年级", Age: 24},
		{Name: "lzw", Grade: "6年级", Age: 34},
	}

	db.Create(&students)
	// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	new_student := [1]Student{{Name: "张三", Grade: "三年级", Age: 20}}
	db.Create(&new_student)

	// // 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	var youngStudents []Student
	result := db.Where("age > ?", 18).Find(&youngStudents)
	if result.Error != nil {
		fmt.Println("查询错误:", result.Error)
	} else {
		fmt.Printf("找到 %d 个学生大于18岁:\n", len(youngStudents))
		for _, s := range youngStudents {
			fmt.Printf("- %s (年龄: %d) 年级:%s\n", s.Name, s.Age, s.Grade)
		}
	}
	fmt.Printf("find 影响行数 %d\n", result.RowsAffected)

	// // // 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	result = db.Model(&Student{}).Where("name = ?", "张三").Updates(Student{Grade: "四年级"})
	if result.Error != nil {
		fmt.Println("更新错误:", result.Error)
	}
	fmt.Printf("updates 影响行数 %d\n", result.RowsAffected)

	// // // 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	result = db.Where("age < ?", 15).Delete(&Student{})
	if result.Error != nil {
		fmt.Println("更新错误:", result.Error)
	}
	fmt.Printf("delete 影响行数 %d\n", result.RowsAffected)

}
