package main

import (
	"fmt"
	"time"
)

func main() {

	//在多个管道中只要有一个操作成功就执行相应逻辑
	channelA := make(chan int)
	channelB := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		close(channelA)
	}()

	select {
	case v, ok := <-channelA:
		fmt.Println("a", v, ok)
	case v, ok := <-channelB:
		fmt.Println("b", v, ok)
	default:
		fmt.Println("default")
	}
}
