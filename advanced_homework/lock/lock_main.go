// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var mu sync.Mutex     // 互斥锁
var wg sync.WaitGroup // 同时管理多个tasks
var lock_shared_counter_number int
var unlock_shared_counter_number int32

func Lock_Shared_Counter_Wirte(num int) {
	defer wg.Done()
	mu.Lock()
	for i := 0; i < num; i++ {
		lock_shared_counter_number += i
	}
	mu.Unlock()
}

// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。

// 普通操作可能被重排序：
// 1. 写buffer
// 2. 其他操作
// 3. 实际写入内存（顺序不确定）

// 原子操作：
// LOCK XADD       ← 内存屏障在这里
// 保证操作前的写入对操作后可见

// 使用建议：

// 简单计数器 → 原子操作

// 单个标志位 → 原子操作

// 复杂数据结构/多个变量 → 互斥锁

// 读多写少 → sync.RWMutex

// 性能关键路径 → 优先考虑原子操作

// 记住：原子操作不是万能的，它只适用于简单的读取-修改-写入操作。复杂的业务逻辑仍然需要锁来保护。
func unLock_Shared_Counter_Wirte(num int) {
	defer wg.Done()
	for i := 0; i < num; i++ {
		atomic.AddInt32(&unlock_shared_counter_number, int32(i))
	}
}
func main() {
	test_count := 0
	for i := 0; i < 1000; i++ {
		test_count += i
	}
	fmt.Printf("自增1000次=%d * 10 = %d\n", test_count, 10*test_count)
	// 测试代码
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Lock_Shared_Counter_Wirte(1000)

	}
	wg.Wait()

	fmt.Println("10个 go lock 线程运行任务完成", lock_shared_counter_number)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go unLock_Shared_Counter_Wirte(1000)

	}
	wg.Wait()
	fmt.Println("10个 go unlock 线程运行任务完成", unlock_shared_counter_number)

}
