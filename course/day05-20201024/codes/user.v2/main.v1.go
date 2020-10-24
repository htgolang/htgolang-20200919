package main

import "fmt"

func main() {
END:
	for {
		var text string
		fmt.Print("请输入指令: ")
		fmt.Scan(&text)

		switch text {
		case "add":
			fmt.Println("添加")
		case "modify":
			fmt.Println("修改")
		case "delete":
			fmt.Println("删除")
		case "query":
			fmt.Println("查询")
		case "exit":
			fmt.Println("退出")
			break END
		default:
			fmt.Println("输入指令错误")
		}
	}
}
