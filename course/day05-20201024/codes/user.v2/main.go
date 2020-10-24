package main

import "fmt"

func addUser() {
	fmt.Println("addUser")
}

func modifyUser() {

	fmt.Println("modifyUser")
}

func deleteUser() {
	fmt.Println("deleteUser")
}

func queryUser() {
	fmt.Println("queryUser")
}

func main() {
	operates := map[string]func(){
		"add":    addUser,
		"modify": modifyUser,
		"delete": deleteUser,
		"query":  queryUser,
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
			op()
		} else {
			fmt.Println("指令错误")
		}

	}
}
