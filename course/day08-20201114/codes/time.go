package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(<-time.After(3 * time.Second))
	for now := range time.Tick(3 * time.Second) {
		fmt.Println(now)
	}
}
