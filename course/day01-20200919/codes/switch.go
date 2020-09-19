package main

import "fmt"

func main() {
	fmt.Println("老婆的想法：")
	fmt.Println("10个包子")
	// 判断 是否有卖西瓜的 (控制台输入y)
	// 有 买一个西瓜

	var text string
	fmt.Print("有卖西瓜的吗:")
	fmt.Scan(&text)
	switch text {
	case "y":
		fmt.Println("买一个西瓜")
	}
}
