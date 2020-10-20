package main

import "fmt"

func main() {
	var zuo int = 1
	//zuo是乘法口诀左边数值，右是乘法口诀右边数值
	for zuo < 10 {
		var you int = 1
		for you <= zuo {
			fmt.Printf("%d * %d = %d ", you, zuo, zuo*you)
			you++
		}
		zuo++
		fmt.Println(" ")
	}
}
