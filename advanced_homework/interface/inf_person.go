package main

import "fmt"

type Person struct {
	Name string
	Age  string
}

type Employee struct {
	Person
	EmployeeID string
}

// 单个Employee的打印方法
func (c *Employee) PrintInfo() {
	fmt.Printf("雇员的个人信息 ID:%s name:%s age:%s\n",
		c.EmployeeID, c.Name, c.Age)
}

// 多个Employee的打印方法
type EmployeeList []Employee

//func (emps *[]Employee) PrintInfo_all() {
/* 为什么不能够这样声明？
*避免全局污染
*如果Go允许为 []Employee 直接添加方法，那么：
*所有 []Employee 都会有这个方法
*不同库可能定义同名方法，产生冲突
*无法区分"普通Person列表"和"有特殊功能的Person列表"
 */

func (emps *EmployeeList) PrintInfo_all() {
	// 关键：使用 *emps 解引用
	for i, emp := range *emps {
		fmt.Printf("%d. ", i+1)
		emp.PrintInfo()
	}
}
