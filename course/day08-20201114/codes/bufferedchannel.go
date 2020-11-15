package main

import "fmt"

func main() {
	var channel chan int = make(chan int, 3)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		// close(channel)
	}()

	// for num := range channel {
	// 	fmt.Println(num)
	// }
	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}

}
