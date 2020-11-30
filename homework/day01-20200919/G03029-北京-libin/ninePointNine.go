package main

import (
	"fmt"
)

//nine nine table
func main() {
	for i := 1; i <= 9; i++ {
		for x := 1; x <= i; x++ {
			fmt.Printf("%d * %d = %d\t", i, x, i*x)
		}
		fmt.Printf("\n")
	}
}
