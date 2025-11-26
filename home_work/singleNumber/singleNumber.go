package main

import "fmt"

func singleNumber(nums []int) int {
	for num_index, num_value := range nums {
		same := false
		for num_next_index, num_next_value := range nums {
			if num_index == num_next_index {
				continue
			}
			if num_value == num_next_value {
				same = true
				break
			}
		}
		if !same {
			return num_value
		}
	}
	return -1
}

func main() {
	// 测试代码
	nums := []int{2, 2, 1}
	fmt.Println("只出现一次的数字是:", singleNumber(nums))

	nums2 := []int{4, 1, 2, 1, 2}
	fmt.Println("只出现一次的数字是:", singleNumber(nums2))
}
