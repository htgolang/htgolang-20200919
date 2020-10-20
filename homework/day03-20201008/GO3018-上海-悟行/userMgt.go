package main

import (
	"fmt"
	"strconv"
	"strings"
)
// ID,name,contact,address
// []map[string][string]

const (
	name      = "name"
	contact = "contact"
	address   = "address"
)

var users = []map[string]string{
	{"id": "1", "name": "tony", "contact": "13512345679", "address": "peking"},
	{"id": "2", "name": "wesley", "contact": "13412345679", "address": "Shanghai"},
	{"id": "3", "name": "elsa", "contact": "13312345679", "address": "hangzhou"},
	{"id": "4", "name": "anna", "contact": "13212345679", "address": "chongqing"},
}

var userInfo = fmt.Sprintf ("%3s %20s %40s %40s", "ID", "name", contact, address)
//用户输入
func input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return strings.TrimSpace(text)
}

//打印用户
func printUser(user map[string]string) {
	fmt.Println(strings.Repeat("#", len(userInfo)))
	fmt.Printf("%3s %20s %40s %40s\n", user["id"], user[name], user[contact], user[address])
}

//添加用户
func genId() int {
	// 生成最大的id
	var rt int
	for _, user := range users {
		userId, _ := strconv.Atoi(user["id"])
		if rt < userId {
			rt = userId
		}
	}
	return rt + 1
}

func addUser() map[string]string{
	user := make(map[string]string)
	user["id"] = strconv.Itoa(genId())
	user["name"] = ""
	user["contact"] = ""
	user["address"] = ""
	return user
}

func add() {
	user := addUser()
	fmt.Println("请输入用户信息:")
	for {
		tempName := input("用户名:")
		if verify_user(tempName) {
			user[name] = tempName
			break
		} else {
			fmt.Println("该用户已存在")
		}
	}
	user[contact] = input("联系方式:")
	user[address] = input("通讯地址:")
	users =append(users, user)
	printUser(user)
}

//查询用户
func query() {
	s := input("请输入查询的用户：")
	fmt.Println(userInfo)
	for _, user := range users {
		if s == "all" || strings.Contains(user[name], s) {
			printUser(user)
		} else {
			fmt.Println("该用户不存在")
			break
		}
	}
}

//修改用户
func modify() {
	s := input("请输入要修改的用户ID:")
	for _, user := range users {
		if s == user["id"] {
			printUser(user)
			switch input("是否确认修改(yes/y):") {
			case "yes", "y":
				for {
					tempName := input("请输入新用户名:")
					if verify_user(tempName) {
						user[name] = tempName
						fmt.Println("用户名已修改")
						printUser(user)
						switch input("是否需要修改联系方式(yes/y):") {
						case "yes", "y":
							user[contact] = input("请输入新的联系方式:")
							fmt.Println("联系方式已修改")
							printUser(user)
						default:
							break
						}	
						switch input("是否需要修改通讯地址(yes/y):") {
						case "yes", "y":
							user[address] = input("请输入新通讯地址:")
							fmt.Println("通讯地址已修改")
							printUser(user)
						default:
							break
						}					
						break
					} else {
						fmt.Println("用户已存在！")
					}
				}
			default:
				break
			}
		}
	}
	
}

//删除用户
func delete() {
	s := input("请输入要删除用户的ID:")
	for index, user := range users {
		if s == user["id"] {
			printUser(user)
			switch input("是否确认删除(yes/y)") {
			case "yes", "y":
				copy(users[index:], users[index+1:])
				newUsers := users[:len(users) - 1]
				for _, user := range newUsers {
					printUser(user)
				}
			default:
				break
			}			
		}
	} 
}

//验证用户是否存在
func verify_user(inputName string) bool {
	for _, user := range users {
		if inputName == user[name] {
			return false
		}
	}
	return true
}

func main() {
	menus := map[string]func() {
		"add":		add,
		"query":	query,
		"modify":	modify,
		"delete":	delete,
	}
	for {
		text := input("请输入您的操作[add/query/modify/delete/exit]:")
		if text == "exit" {
			break
		}
		if menu, ok := menus[text]; ok {
			menu()
		} else {
			fmt.Println("输入的指令错误。")
		}
	}
}