package main

import "fmt"

func main() {
	// 控制台输入一个成绩
	// >= 90 =>A
	// >= 80 => B
	// >= 70 => C
	// >= 60 => D
	// 其他 => E
	var score int
	fmt.Print("请输入成绩：")
	fmt.Scan(&score)

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}
