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

	// 1. 创建调度器实例
	scheduler := NewScheduler()

	// 2. 创建并添加任务

	// 任务1：模拟网络请求
	task1 := Task{
		ID: "网络请求",
		Func: func() {
			time.Sleep(1 * time.Second)
			fmt.Println("任务1: 网络请求完成")
		},
	}
	scheduler.AddTask(task1)

	// 任务2：模拟数据库查询
	task2 := Task{
		ID: "数据库查询",
		Func: func() {
			time.Sleep(2 * time.Second)
			fmt.Println("任务2: 数据库查询完成")
		},
	}
	scheduler.AddTask(task2)

	// 任务3：模拟文件操作
	task3 := Task{
		ID: "文件操作",
		Func: func() {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("任务3: 文件操作完成")
		},
	}
	scheduler.AddTask(task3)

	// 3. 执行所有任务
	fmt.Println("开始执行任务...")
	scheduler.Run()

	// 4. 打印结果
	fmt.Println("\n=== 任务执行结果 ===")
	for _, result := range scheduler.results {
		fmt.Printf("任务: %s, 耗时: %v\n", result.TaskID, result.Elapsed)
	}
}
