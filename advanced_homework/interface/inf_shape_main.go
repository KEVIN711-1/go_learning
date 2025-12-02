// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，
// 创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。

package main

import (
	"fmt"
	"math"
)

// 图形
type Shape interface {
	Area() float64      //面积
	Perimeter() float64 // 周长
}

// 矩形
type Rectangle struct {
	length float64
	width  float64
}

func (rectangle *Rectangle) Area() float64 {
	fmt.Printf("计算矩形的面积：%.2f\n", rectangle.length*rectangle.width)
	return rectangle.length * rectangle.width
}

func (rectangle *Rectangle) Perimeter() float64 {
	fmt.Printf("计算矩形的周长：%.2f\n", 2*(rectangle.length+rectangle.width))
	return 2 * (rectangle.length + rectangle.width)
}

// 圆形
type Circle struct {
	radius float64
}

func (circle *Circle) Area() float64 {
	fmt.Printf("计算圆形的面积：%.2f\n", math.Pi*circle.radius*circle.radius)
	return math.Pi * circle.radius * circle.radius
}

func (circle *Circle) Perimeter() float64 {
	fmt.Printf("计算圆形的周长：%.2f\n", 2*(math.Pi*circle.radius))
	return 2 * (math.Pi * circle.radius)
}

func get_shape_area(p Shape) float64 {
	return p.Area()
}

func get_shape_perimeter(p Shape) float64 {
	return p.Perimeter()
}

func main() {
	Rectangle := &Rectangle{length: 20, width: 10}
	Circle := &Circle{radius: 5}

	fmt.Printf("-------- main p Shape 1 --------  \n")
	//Circle
	Circle_area := get_shape_area(Circle)
	Circle_perimeter := get_shape_perimeter(Circle)
	fmt.Printf("main 计算圆形的面积 %.2f 计算圆形的周长：%.2f\n", Circle_area, Circle_perimeter)

	//Rectangle
	Rectangle_area := get_shape_area(Rectangle)
	Rectangle_perimeter := get_shape_perimeter(Rectangle)
	fmt.Printf("main p Shape %.2f 计算矩形的周长：%.2f\n", Rectangle_area, Rectangle_perimeter)

	fmt.Printf("-------- main p Shape 2 --------  \n")
	//Circle
	Circle_area = Circle.Area()
	Circle_perimeter = Circle.Perimeter()
	fmt.Printf("main 计算圆形的面积 %.2f 计算圆形的周长：%.2f\n", Circle_area, Circle_perimeter)

	//Rectangle
	Rectangle_area = Rectangle.Area()
	Rectangle_perimeter = Rectangle.Perimeter()
	fmt.Printf("main p Shape %.2f 计算矩形的周长：%.2f\n", Rectangle_area, Rectangle_perimeter)

	// 单个Employee
	per_employee := Employee{
		Person: Person{
			Name: "kfzhu",
			Age:  "24",
		},
		EmployeeID: "10",
	}
	fmt.Println("=== 单个雇员 ===")
	per_employee.PrintInfo()

	// Employee切片指针
	per_employees := EmployeeList{
		{
			Person: Person{
				Name: "张三",
				Age:  "24",
			},
			EmployeeID: "10",
		},
		{
			Person: Person{
				Name: "李四",
				Age:  "25",
			},
			EmployeeID: "11",
		},
		{
			Person: Person{
				Name: "王五",
				Age:  "26",
			},
			EmployeeID: "12",
		},
		{
			Person: Person{
				Name: "赵六",
				Age:  "27",
			},
			EmployeeID: "13",
		},
	}

	fmt.Println("\n=== 所有雇员 ===")
	per_employees.PrintInfo_all()
}
