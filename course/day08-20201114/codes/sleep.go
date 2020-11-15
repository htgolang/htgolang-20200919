package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	time.Sleep(time.Second * 3) // 等待3s
	fmt.Println(time.Now())
}
