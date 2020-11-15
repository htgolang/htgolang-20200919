package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 10 个例程分别给salary + 10 1000
// 10 个例程分别给salary - 10 1000
func main() {

	var salary int = 0

	var wg sync.WaitGroup

	// 定义锁
	var locker sync.Mutex

	fmt.Println("start")
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				locker.Lock() // 加锁
				salary += 10  //? n步
				//                      salary 0
				// 内存 -> 寄存器A(1)        0
				// 寄存器A +10 (3)			10
				// 寄存器A -> 内存 (4)		 10
				locker.Unlock() //释放锁
				runtime.Gosched()
			}
		}()
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				locker.Lock()
				salary -= 10 //?
				//                      salary 0
				// 内存 -> 寄存器B(2)		0
				// 寄存器B -10(5)			-10
				// 寄存器B -> 内存(6)		-10
				locker.Unlock()
				runtime.Gosched()
			}
		}()
	}

	wg.Wait()
	fmt.Println("end:", salary)

}
