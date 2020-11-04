package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %-4v", i, j, i * j )
		}
		fmt.Println()
	}
}