package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d ", j, i, i*j)
		}
		fmt.Printf("\n")
	}
}
