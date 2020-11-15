package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// 10 个例程分别给salary + 10 1000
// 10 个例程分别给salary - 10 1000
// sync/atomic
// Add系列
// Load系列
// Store系列
// Swap系列
// CompareAndSwap系列

func main() {

	var salary int32 = 0

	var wg sync.WaitGroup

	fmt.Println("start")
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				// salary += 10 //? n步
				atomic.AddInt32(&salary, 10)
				//                      salary 0
				// 内存 -> 寄存器A(1)        0
				// 寄存器A +10 (3)			10
				// 寄存器A -> 内存 (4)		 10
				runtime.Gosched()
			}
		}()
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				// salary -= 10 //?
				atomic.AddInt32(&salary, -10)
				//                      salary 0
				// 内存 -> 寄存器B(2)		0
				// 寄存器B -10(5)			-10
				// 寄存器B -> 内存(6)		-10
				runtime.Gosched()
			}
		}()
	}

	wg.Wait()
	fmt.Println("end:", salary)

}
