package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan []int, 1) // type
	num := []int{1, 2, 3}
	channel <- num

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		a := <-channel
		a[0] = 100
		fmt.Println(a) // 修改a, 有可能影响num吗
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(num)

	/*
		channel
		int 无影响
		array 无影响
		slice/map 可能会有影响
		channel <- struct{}{} // 通知
		<- 0
		<- ""
	*/
}
