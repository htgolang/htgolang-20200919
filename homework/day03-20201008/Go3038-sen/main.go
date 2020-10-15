package main

import (
	"day03/funcs"
	"fmt"
)

//
var Users = []map[string]string{}

func main()  {
	fmt.Println("=========== User Manager ==========")
	fmt.Printf(`
		1. 添加用户
		2. 修改用户
		3. 查询
		4. 删除用户
		5. 退出

`)
	options := map[string]func(){
		"1": funcs.Add,
		"2": funcs.Modify,
		"3": funcs.Query,
		"4": funcs.Del,
		"5": funcs.Exit,
	}

	for {
		userInput := funcs.Input("请输入选项：")
		if options, ok := options[userInput]; ok {
			options()
		}else {
			fmt.Println("没有这个选项，请重新输入")
		}
	}

	//for {
	//	switch (funcs.Input("请选择操作（add/query/exit/...）:")) {
	//	case "add":
	//		funcs.Add()
	//	case "query":
	//		funcs.Query()
	//	case "delete":
	//		funcs.DeleteUser()
	//	case "exit":
	//		return
	//	default:
	//		fmt.Println("错误的指令")
	//	}
	//}

}
