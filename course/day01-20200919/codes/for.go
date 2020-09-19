package main

import "fmt"

func main() {
	// 1 + 2 + ... + 100
	sum := 0

	for idx := 1; idx <= 100; idx++ {
		sum += idx
	}

	fmt.Println(sum)
}
