package main

import "fmt"

// 只写
func in(channel chan<- int) {
	channel <- 1
	channel <- 2
	// fmt.Println(<-channel)
	channel <- 3
}

// 只读
func out(channel <-chan int) {
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	// channel <- 4
}

func main() {
	var channel chan int = make(chan int, 3)
	var rchannel <-chan int = channel
	var wchannel chan<- int = channel
	in(wchannel)
	out(rchannel)

	fmt.Println(<-channel)

}
