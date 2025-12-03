// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，
// 另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。

package main

import (
	"fmt"
	"sync"
)

// 只发送通道入参定义
// func channel_send(c <- chan int, sum int) {
var wg sync.WaitGroup

func channel_send(c chan int, sum int) {
	defer wg.Done()
	for i := 0; i < sum; i++ {
		c <- i
	}
	close(c) // 发送完成后关闭通道, 否则会死锁
}

// 只接受通道入参定义
// func channel_read(c chan<- int) {
func channel_read(c chan int) {
	defer wg.Done()

	for v := range c {
		fmt.Printf("channel value:%d\n", v)
	}
}

func main() {
	// 测试代码
	trrans_cha := make(chan int, 10)

	wg.Add(1)
	go channel_send(trrans_cha, 100)

	wg.Add(1)
	go channel_read(trrans_cha)

	//为什么需要增加wg sync.WaitGrroup 的处理函数？
	//这是因为主 goroutine 提前退出了，没有等待其他 goroutine 执行完成。
	//需要等待两个子go 运行完才行
	wg.Wait() // 等待所有goroutine完成，所有wg计数为0退出
	fmt.Println("所有任务完成")
}
