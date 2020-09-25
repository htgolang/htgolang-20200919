package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v X %v = %v\t", j, i, i*j)
		}
		fmt.Printf("\n")
	}
}
