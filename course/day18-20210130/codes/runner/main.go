package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	closeChan := make(chan struct{})

	go func() {
		ticker := time.NewTicker(time.Second * 3)
	STOP:
		for {
			select {
			case <-closeChan:
				fmt.Println("close")
				break STOP
			case now := <-ticker.C:
				fmt.Println(now)

			}
		}
		fmt.Println("over")
		wg.Done()
	}()

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
	close(closeChan)
	fmt.Println("interrupt")

	wg.Wait()
}
