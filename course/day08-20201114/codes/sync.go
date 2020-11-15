package main

import (
	"fmt"
	"sync"
)

func task() {
	fmt.Println("task")
}

func taskA() {
	fmt.Println("taska")
}

func main() {
	// sync
	// sync.WaitGroup 计数信号量
	// sync.Mutex 锁 互斥锁
	// 共享数据 读
	// 共享数据 写
	// sync.RWMutex 读写锁
	// sync.Cond 条件锁
	// 多个例程, 某个执行检查是否满足条件，不满足等待 Wait
	// 其他例程, 当可能产生等待例程条件重新满足, 通知等待例程 Signal/Boardcase
	// sync.Map
	// map + sync.RWMutex

	// once := &sync.Once{}
	// once.Do(task) // 函数值运行一次
	// once.Do(taskA)
	// once.Do(task)

	// sync.Pool 线程池， 连接池，...
	// 对象池，从池中获取对象，当池中无可用对象，创建并返回
	// 当使用完成会放入池中
	intPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("new")
			return 1
		},
	}

	v := intPool.Get() // 需要创建 -> New
	// v断言某类型，执行
	fmt.Println(v)
	// 使用完成放入池中
	intPool.Put(v)      // 1个对象
	v1 := intPool.Get() //
	v2 := intPool.Get() // 需要创建 -> New
	fmt.Println(v1, v2)
}
