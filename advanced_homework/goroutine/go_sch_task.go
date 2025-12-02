// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），
// 并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。

package main

import (
	"sync"
	"time"
)

type Task struct {
	ID   string
	Func func()
}

type TaskResult struct {
	TaskID  string        // 任务ID
	Elapsed time.Duration // 执行耗时
}

type Scheduler struct {
	tasks   []Task         // 待执行任务列表
	results []TaskResult   // 任务执行结果
	mu      sync.Mutex     // 互斥锁
	wg      sync.WaitGroup // 同时管理多个tasks
}

// AddTask 向调度器添加任务
func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

// Run 执行所有任务并统计时间
func (s *Scheduler) Run() []TaskResult {
	// 遍历任务，启动协程执行
	for _, task := range s.tasks {
		s.wg.Add(1) //（计数+1）
		// 启动协程执行单个任务
		go func(t Task) {
			//defer 任务完成后 调用，防止goroutine 中途调用 return
			defer s.wg.Done() // 减少计数 （计数-1）

			// 记录任务开始时间
			start := time.Now()
			// 执行任务
			t.Func()
			// 计算执行耗时
			elapsed := time.Since(start)

			// 并发安全地写入结果
			/*
			* 只有一个goroutine 能获得锁
			* 其他两个等待（排队）
			* 保证同一时间只有一个goroutine修改results
			 */
			s.mu.Lock()
			s.results = append(s.results, TaskResult{
				TaskID:  t.ID,
				Elapsed: elapsed,
			})
			s.mu.Unlock()
		}(task) // 传入当前任务（避免循环变量捕获问题）
	}

	// 等待所有任务执行完成
	s.wg.Wait() //（阻塞直到计数为0）
	return s.results
}

// 构造函数（帮助你创建Scheduler实例）
func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks:   make([]Task, 0),
		results: make([]TaskResult, 0),
	}
}
