package funcs

import (
	"fmt"
	"homework/pkg/models"
)

func userQuery() {
	//提供4种查询方式
	var userChoice string
	fmt.Println(`请输入要查询的类型：
	1.根据ID查询
	2.根据Name查询
	3.根据Contact查询
	4.根据Address查询
	`)
	fmt.Scan(&userChoice)
	switch userChoice {
	case "1":
		userQueryExec("ID")
	case "2":
		userQueryExec("Name")
	case "3":
		userQueryExec("Contact")
	case "4":
		userQueryExec("Address")	
	default:
		fmt.Println("选择有误，退出！")
	}
}

func userQueryExec(elemet string) {
	var userInput,state string
	users := make([]map[string]string,0)
	fmt.Printf("请输入要查询%s的信息：\n",elemet)
	fmt.Scan(&userInput)
	for _,v:=range models.Users {
		if v[elemet] == userInput {
			users = append(users,v)
			state = "True"
		}
	}
	if state == "True" {
		usersList(&users)
	} else {
		fmt.Printf("未找到%s为%s的数据。\n",elemet,userInput)
	}
}