package main

import "fmt"

func main() {
	// 1 + 2 + 100

	sum := 0
	idx := 1

START:
	if idx > 100 {
		goto END
	}

	sum += idx
	idx += 1
	goto START

END:
	fmt.Println(sum)
}
