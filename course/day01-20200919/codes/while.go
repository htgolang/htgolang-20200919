package main

import "fmt"

func main() {
	// 1 + 2 ... + 10
	sum := 0
	idx := 1

	for idx <= 100 {
		sum += idx
		idx++
	}

	fmt.Println(sum)
}
