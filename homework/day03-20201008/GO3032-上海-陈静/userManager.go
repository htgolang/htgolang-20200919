package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	id    string
	name  string
	phone string
	addr  string
	users = []map[string]string{}
)

func addUser() {
	userInfo := map[string]string{
		"id":    "",
		"name":  "",
		"phone": "",
		"addr":  "",
	}

	fmt.Print("姓名：")
	fmt.Scan(&name)
	fmt.Print("\n联系方式(电话)：")
	fmt.Scan(&phone)
	fmt.Print("\n通信地址：")
	fmt.Scan(&addr)
	fmt.Println(len(users))

	if len(users) == 0 {
		id = "1"
	} else {
		id = strconv.Itoa(len(users) + 1)
	}

	userInfo["id"] = id
	userInfo["name"] = name
	userInfo["phone"] = phone
	userInfo["addr"] = addr

	users = append(users, userInfo)

	fmt.Println(users)
}

func delUser() {
	if len(users) == 0 {
		fmt.Println("空用户!!!")
		return
	}

	var isDelete string

	fmt.Print("id：")
	fmt.Scan(&id)

	for i, v := range users {
		if v["id"] == id {
			fmt.Println(v)
			fmt.Print("请输入y删除：")
			fmt.Scan(&isDelete)

			if isDelete == "y" && i == 0 {
				users = users[1:]
			} else if isDelete == "y" && i == len(users)-1 {
				users = users[:len(users)-1]
			} else if isDelete == "y" {
				users = append(users[:i], users[i+1:]...)
			}
		}
	}

	fmt.Println(users)
}

func modifyUser() {
	var isModify string

	fmt.Print("id：")
	fmt.Scan(&id)

	for _, v := range users {
		if v["id"] == id {
			fmt.Println(v)
			fmt.Print("请输入y修改：")
			fmt.Scan(&isModify)

			if isModify == "y" {
				fmt.Print("姓名：")
				fmt.Scan(&name)
				fmt.Print("\n联系方式(电话)：")
				fmt.Scan(&phone)
				fmt.Print("\n通信地址：")
				fmt.Scan(&addr)

				v["name"] = name
				v["phone"] = phone
				v["addr"] = addr

				fmt.Println(v)
			} else {
				fmt.Println("取消修改")
			}

		}
	}

	fmt.Println(users)
}

func queryUser() {
	if len(users) == 0 {
		fmt.Println("空用户!!!")
		return
	}

	var info string
	var flag bool
	fmt.Print("请输入要查找的信息：")
	fmt.Scan(&info)
	for i, v := range users {
		flag = false
		for _, in := range v {
			if strings.Contains(in, info) && flag == false {
				flag = true
				fmt.Println(users[i])
			}
		}
	}
}

func main() {
	for {
		fmt.Println(`
---用户管理---
1、添加用户
2、删除用户
3、修改用户
4、查找用户
5、退出(q键)`)

		fmt.Printf("%v", users)
		var userchoice string
		fmt.Print("请选择：")
		fmt.Scan(&userchoice)
		switch {
		case userchoice == "1":
			addUser()
		case userchoice == "2":
			delUser()
		case userchoice == "3":
			modifyUser()
		case userchoice == "4":
			queryUser()
		case strings.ToLower(userchoice) == "q":
			os.Exit(0)
		default:
			fmt.Println("请选择(1/2/3/4)")
		}
	}

}
