package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result [2]int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {
				result[0] = i
				result[1] = j
				return result[:]
			}
		}
	}
	// 	return result        // 错误！返回 [2]int，但函数需要 []int
	//  return result[:]     // 正确！返回 []int（切片）
	return result[:]
}

func main() {
	// 测试代码
	nums := []int{2, 7, 11, 15}
	fmt.Println("两数之和:", twoSum(nums, 13))

	nums = []int{3, 2, 4}
	fmt.Println("两数之和:", twoSum(nums, 7))

	nums = []int{3, 3}
	fmt.Println("两数之和:", twoSum(nums, 6))
}
