package main

import (
	"fmt"
	"math"
	"strconv"
)

func plusOne(digits []int) []int {
	var big_num, zero_count int

	for i := 0; i < len(digits); i++ {
		zero_count = len(digits) - i - 1

		if zero_count == 0 {
			big_num += digits[i]
		} else {
			big_num += digits[i] * int(math.Pow(10, float64(zero_count)))
		}

	}

	big_num++

	num_str := strconv.Itoa(big_num)
	// str_array := []byte(num_str) // 转换为 rune 切片

	plus_one := make([]int, len(num_str))

	for i := 0; i < len(num_str); i++ {
		// digit, _ := strconv.Atoi(string(str_array[i]))
		// plus_one[i] = digit
		plus_one[i] = int(num_str[i] - '0') // 直接字符运算，避免类型转换
	}
	return plus_one
}

func main() {
	plus_one := []int{1, 2, 3}
	fmt.Println("plusone:", plusOne(plus_one))

	plus_one = []int{4, 3, 2, 2}
	fmt.Println("plusone:", plusOne(plus_one))

	plus_one = []int{9}
	fmt.Println("plusone:", plusOne(plus_one))
}
