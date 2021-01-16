package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 管道
	fmt.Println("pid:", os.Getpid())

	interrupt := make(chan os.Signal, 1)
	reload := make(chan os.Signal, 1)
	// 监听系统信号
	signal.Notify(interrupt, syscall.SIGINT)
	signal.Notify(reload, syscall.SIGHUP)

	// 等待信号
INTERRUPT:
	for {
		select {
		case <-interrupt:
			break INTERRUPT
		case <-reload:
			fmt.Println("reload")
		}
	}

}
