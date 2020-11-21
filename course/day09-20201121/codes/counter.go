package main

import "time"

func main() {
	// 计数停止例程的数量
	var stopCounter int = 0 // 锁的功能

	const counter = 3
	// 3个例程
	for i := 0; i < counter; i++ {
		go func() {
			stopCounter++
		}()
	}

	//如何等待stopCounter是3
	for stopCounter != counter {
		time.Sleep(3 * time.Second)
	}
}
