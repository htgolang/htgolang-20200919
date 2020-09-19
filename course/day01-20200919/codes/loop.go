package main

import "fmt"

func main() {
	// 1 + 2 ... + 100
	sum := 0
	idx := 1

	for {
		sum += idx
		idx++
		if idx > 100 {
			break
		}
	}

	fmt.Println(sum)
}
