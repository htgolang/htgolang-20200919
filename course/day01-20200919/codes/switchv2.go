package main

import "fmt"

func main() {
	fmt.Println("老公的想法：")
	// 判断 是否有卖西瓜的 (控制台输入y)
	// 有 买1个包子
	// 否则买10个包子

	var text string
	fmt.Print("有卖西瓜的吗:")
	fmt.Scan(&text)
	switch text {
	case "y":
		fmt.Println("买一个包子")
	default:
		fmt.Println("买10个包子")
	}
}
