package main

import (
	"fmt"
	"runtime"
	"sync"
)

func taskA(wg *sync.WaitGroup) {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		// time.Sleep(time.Microsecond * 1000)
		runtime.Gosched() // 让出调度
	}

	// 等待信号量 -1
	wg.Done()
}

func taskB(wg *sync.WaitGroup) {
	// 等待信号量 -1
	defer wg.Done()

	for i := 'A'; i <= 'Z'; i++ {
		fmt.Printf("%c\n", i)
		// time.Sleep(time.Microsecond * 1000)
		runtime.Gosched() // 让出调度
	}

}

func taskM() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c\n", i)
		runtime.Gosched()
	}
}

func main() {
	fmt.Println("start")
	// taskA， taskB先执行，不能预知
	// A 1 2 ... 10 B ... Z

	// 计数信号量
	// 启动例程的之前, +1
	// 当例程执行结束时, -1

	// var wg sync.WaitGroup
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go taskA(wg) // 启动一个例程
	go taskB(wg) // 启动一个例程
	// time.Sleep(5 * time.Second)
	// taskM()

	wg.Wait()
	fmt.Println("end")

	// goroutine => 例程

	// main例程 => 主例程
	// go 工作例程
	// 主例程 不等待工作例程执行结束
}
