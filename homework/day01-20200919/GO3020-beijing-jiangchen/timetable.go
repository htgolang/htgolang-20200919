package main

import (
	"fmt"
)

//Timetable ...
func Timetable() {
	fmt.Println("========打印九九乘法表========")
	for y := 1; y <= 9; y++ {
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d ", x, y, x*y)
		}
		fmt.Println()
	}
	fmt.Println("========打印结束========")
}
