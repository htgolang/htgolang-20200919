package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		ticker := time.NewTicker(3 * time.Second)
	STOP:
		for {
			select {
			case <-ctx.Done():
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
	// cancel() 工作例程退出了但是主例程等待信号
	<-interrupt
	cancel()
	fmt.Println("interrupt")
	wg.Wait()
	// cancel()
}
