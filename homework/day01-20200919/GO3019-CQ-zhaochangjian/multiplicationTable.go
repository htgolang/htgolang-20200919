package main

import "fmt"

func main() {
	// 打印九九乘法口诀表
	// 方法一
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j < i {
				fmt.Print(j, "*", i, "=", j*i, " ")
			} else if j == i {
				fmt.Print(j, "*", i, "=", j*i, "\n")
			}
		}
	}
	fmt.Println("=============")

	// 方法二，可对齐打印方案
	for i := 1; i <= 9; i++ {
		oneLine := ""
		for j := 1; j <= i; j++ {
			oneLine += fmt.Sprintf("%-2d", j) + "*" + fmt.Sprintf("%2d", i) + " = " + fmt.Sprintf("%-2d", j*i) + "  "

		}
		fmt.Println(oneLine)
	}
	fmt.Println("=============")

	// 方法三，可对齐打印方案
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%-2d  ", j, i, j*i)
		}
		fmt.Println()
	}
}

