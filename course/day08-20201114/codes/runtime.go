package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for {

		}
	}()
	// 开发终端 对CPU要求有限制 4核 1
	// runtime.Gosched()
	runtime.GOMAXPROCS(1) // 设置使用CPU核数
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

}
