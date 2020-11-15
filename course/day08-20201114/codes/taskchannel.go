package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("start")

	channel := make(chan int)
	cnt := 2

	go func() {
		for i := 'A'; i <= 'Z'; i++ {
			fmt.Printf("%c\n", i)
			runtime.Gosched()
		}

		channel <- 0
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Printf("%d\n", i)
			runtime.Gosched()
		}

		channel <- 0
	}()

	// time.Sleep(3 * time.Second)
	for i := 0; i < cnt; i++ {
		<-channel
	}
	fmt.Println("end")
}
