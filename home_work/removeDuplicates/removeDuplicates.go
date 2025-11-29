package main

import "fmt"

func removeDuplicates(nums []int) int {
	var removeDuplicates_count, removeDuplicates_num int
	dup_ums := []int{}
	var dup_count int
	is_dup := false

	for i := 0; i < len(nums); i++ {
		removeDuplicates_num = nums[i]
		is_dup = false

		for j := i + 1; j < len(nums); j++ {
			if removeDuplicates_num == nums[j] {
				is_dup = true
				break
			}
		}

		if is_dup == true {
			dup_count++
		} else {
			dup_ums = append(dup_ums, removeDuplicates_num)
		}
	}

	removeDuplicates_count = len(nums) - dup_count
	/* nums = dup_ums[:removeDuplicates_count]
	* 切片无法修改外部数组的值
	* 可以使用copy函数修改外部数组的值
	 */
	copy(nums, dup_ums[:removeDuplicates_count])

	/* 可以直接通过数组地址赋值修改外部数组的值*/
	// for k := 0; k < len(dup_ums); k++ {
	// 	nums[k] = dup_ums[k]
	// }

	return removeDuplicates_count
}

func main() {
	int_array := []int{1, 1, 2}
	fmt.Printf("len_length:%d, int_array=%v\n", removeDuplicates(int_array), int_array)

	int_array = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Printf("len_length:%d, int_array=%v\n", removeDuplicates(int_array), int_array)
}
