package main

import "fmt"

func point_num_add(target *int, add_num int) {
	*target += add_num
}

func point_array_multi(target []int, multi_num int) {
	for index, value := range target {
		target[index] = value * multi_num
	}
}

func main() {
	// 测试代码
	num := 10
	fmt.Println("before point_1:", num)

	point_num_add(&num, 10)
	fmt.Println("after point_1:", num)

	num_array := []int{10, 1, 2, 3, 4, 5, 70}
	fmt.Println("before point_2:", num_array)

	point_array_multi(num_array, 2)
	fmt.Println("after point_2:", num_array)

}
