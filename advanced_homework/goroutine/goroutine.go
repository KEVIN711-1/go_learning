package main

import (
	"fmt"
	"time"
)

func goroutine_odd_even_number() {
	go func() {
		num := 10
		for i := 0; i < num; i++ {
			if i%2 == 1 {
				fmt.Println("goroutine 1 print odd number:", i)
			}
		}
	}()

	go func() {
		num := 10
		for i := 0; i < num; i++ {
			if i%2 == 0 {
				fmt.Println("goroutine 2 print even number:", i)
			}
		}
	}()
}

func main() {
	// 测试代码
	goroutine_odd_even_number()
	// 等待goroutine执行完成
	time.Sleep(1 * time.Second)
}
