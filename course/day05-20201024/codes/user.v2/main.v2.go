package main

import "fmt"

func main() {
	operates := map[string]string{
		"add":    "添加",
		"modify": "修改",
		"delete": "删除",
		"query":  "查询",
	}
	for {
		var text string
		fmt.Print("请输入指令: ")
		fmt.Scan(&text)
		if text == "exit" {
			fmt.Println("退出")
			break
		}
		if op, ok := operates[text]; ok {
			fmt.Println(op)
		} else {
			fmt.Println("指令错误")
		}

	}
}
