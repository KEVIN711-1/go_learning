package main

import "fmt"

func point_1(target *int) {
	*target += 10
}

func point_2(target []int) {
	for index, value := range target {
		target[index] = value * 2
	}
}

func main() {
	// 测试代码
	num := 10
	fmt.Println("before point_1:", num)

	point_1(&num)
	fmt.Println("after point_1:", num)

	num_array := []int{10, 1, 2, 3, 4, 5, 70}
	fmt.Println("before point_2:", num_array)

	point_2(num_array)
	fmt.Println("after point_2:", num_array)

}
