package main

import "fmt"

const COUNT int = 10

func main() {
	for i := 1; i <= COUNT; i++ {
		for j := 1; j <= i; j++ {

			fmt.Printf("%2d * %2d = %2d\t", i, j, i*j)
			//Printf  将数据输出到按照格式输出到标准输出流中
			// %nd    表示最小占位n个宽度且右对齐
			//	\t	  相当于tab，缩进
			//	\r	  回车
			//	\n	  换行符
			//	\b	  换成一个黑点
			//	\"	  转义"
			//	\'	  转义'
			//	\\	  转义\
		}
		fmt.Println() //Println：将数据输出到标准输出流中，并添加换行
	}
}