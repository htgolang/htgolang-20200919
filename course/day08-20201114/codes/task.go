package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("start")

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for i := 'A'; i <= 'Z'; i++ {
			fmt.Printf("%c\n", i)
			runtime.Gosched()
		}

		wg.Done()
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Printf("%d\n", i)
			runtime.Gosched()
		}

		wg.Done()
	}()

	// time.Sleep(3 * time.Second)
	wg.Wait()
	fmt.Println("end")
}
