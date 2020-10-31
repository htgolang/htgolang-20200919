package main

import "fmt"

func main() {
	// 语言级别的缓冲
	var num int

	fmt.Println(fmt.Scan(&num))
	fmt.Println(num)

	fmt.Println(fmt.Scan(&num))
	fmt.Println(num)

	fmt.Println(fmt.Scan(&num))
	fmt.Println(num)

	fmt.Println(fmt.Scan(&num))
	fmt.Println(num)

	fmt.Println(fmt.Scan(&num))
	fmt.Println(num)
}
