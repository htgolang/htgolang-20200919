package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// 使用for循环启动3各例程
	fmt.Println("start") //0,1,2=>3,3,3
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("end")
}
