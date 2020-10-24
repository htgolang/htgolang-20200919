package main

import "fmt"

type Counter = int // 别名

func main() {

	var counter Counter

	fmt.Printf("%T, %v\n", counter, counter)
	var num int = 10
	fmt.Println(counter + num)

	var (
		r rune
		b byte
	)
	fmt.Printf("%T, %T", r, b)
}
