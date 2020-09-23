package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		//fmt.Println()
		fmt.Printf("\n")
	}
}

/*
1. 定义变量i为1，i如果小于和等于9，i进行+1
2. 嵌套循环,定义变量j为1，j如果小于或等于i，j进行加1
3. %d打印数值类型整数，打印i至于j值以及i*j的值，\t打印制表符
4. 内循环外进行fmt.Pirntln()打印空值，不过主要是使用Println的换行操作
*/