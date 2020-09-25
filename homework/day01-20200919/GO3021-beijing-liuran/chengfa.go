package main

import (
	"fmt"
)

func main() {
	for a := 1; a <= 9; a++ {
		for b := 1; b <= a; b++ {
			fmt.Printf("%d*%d=%d ", a, b, a*b)
		}
		fmt.Printf("\n")
	}

}
