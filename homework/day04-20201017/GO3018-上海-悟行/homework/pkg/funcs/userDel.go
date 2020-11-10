package funcs

import (
	"fmt"
	"homework/pkg/models"
)

func userDel() {
	/*
	2种删除方式：ID或name，存在则删除，否则提示不存在。
	*/
	var userChoice string
	fmt.Println(`请输入要删除的类型：
	1.根据ID删除
	2.根据Name删除
	`)
	fmt.Scan(&userChoice)
	switch userChoice {
	case "1":
		userDelExec("ID")
	case "2":
		userDelExec("Name")
	default:
		fmt.Println("选择有误，退出！")
	}
}

func userDelExec(elemet string) {
	//根据用户选择的类型来删除
	//userData删除的具体信息；state信息返回状态
	var userInput,state string
	fmt.Printf("请输入要删除%s的信息：\n",elemet)
	fmt.Scan(&userInput)
	for k,v:=range models.Users {
		if v[elemet] == userInput {
			fmt.Println("要删除的数据如下：")
			userIdList(k)
			models.Users = append((models.Users)[:k],(models.Users)[k+1:]...)
			fmt.Printf("%s为%s的数据已删除。\n",elemet,userInput)
			state = "True"
			return
		}
	}
	state = "False"
	if state == "False" {
		fmt.Printf("未找到%s为%s的数据。\n",elemet,userInput)
	}
}