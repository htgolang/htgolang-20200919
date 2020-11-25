package main

import (
	"fmt"
)

func UserMsg() {
	fmt.Printf(`
	-----------------
	1、查看用户
	2、添加用户
	3、删除用户
	4、修改用户
	5、查找用户
	h、查看帮助
	q、退出（exit）
	-----------------
	`)
}

var Users []map[string]string

func Input(s string) string {
	fmt.Printf("%s", s)
	fmt.Scan(&s)
	return s
}

func AddUser() {

	usermap := make(map[string]string)
	usermap["username"] = Input("用户名:")
	usermap["password"] = Input("密码:")
	usermap["tel"] = Input("联系方式:")
	usermap["addr"] = Input("地址:")
	Users = append(Users, usermap)
	fmt.Println(Users)
}

func GetAllUser() {
	fmt.Println(Users)
}

func GetUser() {
	u := Input("查找用户名:")
	for k, v := range Users {
		fmt.Println(k, v)
		if v["username"] == u {
			fmt.Println(Users[k])
		}
	}
}

func DelUser() {
	u := Input("删除用户名：")

	for k, v := range Users {
		fmt.Println(k, v)
		if v["username"] == u {
			fmt.Println(Users[k])
			Users = append(Users[:k], Users[k+1:]...)
		}
	}
	fmt.Println(Users)
}

func ModifUser() {
	usermap := make(map[string]string)
	u := Input("用户名:")
	for k, v := range Users {
		if v["username"] == u {
			usermap["username"] = u
			usermap["password"] = Input("密码:")
			usermap["tel"] = Input("联系方式:")
			usermap["addr"] = Input("地址:")
			Users[k] = usermap
		}
	}
	fmt.Println(Users)

}
func main() {
	// fmt.Println(Input("用户名"))
	UserMsg()
	for {
		// var i string
		i := Input("请输入需要操作的功能：")
		switch i {
		case "1":
			GetAllUser()
		case "2":
			AddUser()
		case "3":
			DelUser()
		case "4":
			ModifUser()
		case "5":
			GetUser()
		case "h":
			UserMsg()
		case "q":
			return
		default:
			UserMsg()
			break
		}
	}

}
