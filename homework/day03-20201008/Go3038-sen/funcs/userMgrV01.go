package funcs

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Users = make([]map[string]string, 0)

const (
	id = "id"
	name = "name"
	contact = "contact"
	address = "address"
)

func isExistName(userName string) bool {
	for _, user := range Users{
		if userName == user[name] {
			return true
		}
	}
	return false
}

func isExistID(userID string) bool {
	for _, user := range Users {
		if userID == user[id]{
			return true
		}
	}
	return false
}

func Input(prompt string) string {
	var text string
	fmt.Println(prompt)
	fmt.Scan(&text)
	return strings.TrimSpace(text)
}

func GenId() int {
	var rt int
	for _, user := range Users {
		userId, _ := strconv.Atoi(user["id"])
		if rt < userId {
			rt = userId
			fmt.Println(rt)
		}
	}
	return rt + 1
}

func printUser(user map[string]string) {
	fmt.Println(strings.Repeat("-", 15))
	fmt.Println("ID:", user[id])
	fmt.Println("Name:", user[name])
	fmt.Println("Contact:", user[contact])
	fmt.Println("Address:", user[address])
	fmt.Println(strings.Repeat("-", 15))
	fmt.Println(Users)
}

func newUser() map[string]string {
	user := make(map[string]string)
	user[id] = strconv.Itoa(GenId())
	user[name] = ""
	user[contact] = ""
	user[address] = ""
	return user
}

func Add()  {
	user := newUser()

	fmt.Println("请输入用户信息")
	user[name] = Input("用户名：")
	if isExistName(user[name]) {
		fmt.Println("输入的用户已存在,请重新输入。")
		return
	}
	user[contact] = Input("联系电话：")
	user[address] = Input("地址：")
	printUser(user)
	Users = append(Users, user)
	fmt.Println("用户创建成功", user)

}

func Del() {
	userID := Input("请输入要删除用户的ID：")
	for idx, user := range Users {
		if isExistID(userID){
			printUser(user)
			fmt.Println(idx, user[id])
			if isDelete := Input("确认删除当前用户信息（yes/y):");
				isDelete == "yes" || isDelete == "y" {
				copy(Users[idx:], Users[idx+1:])
				Users = Users[0:len(Users)-1]
			}
		}else {
			fmt.Println("输入的id不存在")
			return
		}
	}
}

func Modify() {
	id := Input("请输入要编辑的ID：")
	for i, j := range Users {
		if id == j["id"] {
			printUser(j)
			fmt.Println("请输入修改信息：")
			for {
				UserName := Input("用户名称：")
				if !isExistName(UserName) {
					j[name] = UserName
					break
				} else {
					fmt.Println("用户名称已存在， 请重新输入")
				}
			}
			j[contact] = Input("联系方式：")
			j[address] = Input("地址：")
		}
		if i == len(Users) -1{
			fmt.Println("没有这个ID")
		}
	}
}

func Query() {
	que := Input("输入要查询的用户名，查询所有用户输入（all）")

	for _, user := range Users {
		if que == "all" || strings.Contains(user[name], que) {
			printUser(user)
		}
	}
}

func Exit() {
	os.Exit(-1)
}