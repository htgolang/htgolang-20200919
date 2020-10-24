package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

var password = "88cdb337f8c62dc69c1aee4066f80bf5"

var users = []map[string]string{}

// 输入信息
func input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return text
}

// 添加时生成用户ID
func getId() string {
	id := 0
	for _, user := range users {
		if uid, err := strconv.Atoi(user["id"]); err == nil {
			if uid > id {
				id = uid
			}
		}
	}
	return strconv.Itoa(id + 1)
}

// 添加用户
func addUser() {
	user := map[string]string{
		"id":   getId(),
		"name": input("请输入用户名: "),
		"addr": input("请输入联系地址: "),
		"tel":  input("请输入联系方式: "),
	}

	users = append(users, user)
}

// 根据ID查找用户
func findUserById(id string) map[string]string {
	for _, user := range users {
		if user["id"] == id {
			return user
		}
	}
	return nil
}

// 根据ID修改用户数据
func modifyUserById(user map[string]string, id string) {
	for idx, tuser := range users {
		if tuser["id"] == id {
			users[idx] = user
			break
		}
	}
}

func modifyUser() {
	id := input("请输入需要修改的用户ID: ")
	user := findUserById(id)
	if user == nil {
		fmt.Println("用户信息不存在")
	} else {
		fmt.Println("你将要修改的用户信息如下: ")
		fmt.Println(user)
		confirm := input("确定修改吗?(Y/n)")
		if strings.ToLower(confirm) == "y" || strings.ToLower(confirm) == "yes" {
			user := map[string]string{
				"id":   id,
				"name": input("name: "),
				"addr": input("addr: "),
				"tel":  input("tel: "),
			}
			modifyUserById(user, id)
		}
	}
}

// 根据ID删除用户
func deleteUserById(id string) {
	tempUsers := make([]map[string]string, 0, len(users)-1)
	for _, user := range users {
		if user["id"] != id {
			tempUsers = append(tempUsers, user)
		}
	}
	users = tempUsers
}

// 删除用户信息
func deleteUser() {
	id := input("请输入需要删除的用户ID: ")
	user := findUserById(id)
	if user == nil {
		fmt.Println("用户信息不存在")
	} else {
		fmt.Println("你将要删除的用户信息如下: ")
		fmt.Println(user)
		confirm := input("确定删除吗?(Y/n)")
		if strings.ToLower(confirm) == "y" || strings.ToLower(confirm) == "yes" {
			deleteUserById(id)
		}
	}
}

// 过滤用户数据
func filter(user map[string]string, q string) bool {
	return strings.Contains(user["name"], q) ||
		strings.Contains(user["addr"], q) ||
		strings.Contains(user["tel"], q)
}

// 打印用户数据
func printUser(user map[string]string) {
	fmt.Println(user)
}

// 查询用户
func queryUser() {
	q := input("请输入查询信息: ")
	fmt.Println("查询结果")
	for _, user := range users {
		if filter(user, q) {
			printUser(user)
		}
	}
}

func md5text(text string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}

func auth() bool {
	for i := 0; i < 3; i++ {
		if md5text(input("请输入密码: ")) == password {
			return true
		} else {
			fmt.Println("输入密码错误")
		}
	}
	return false
}

func main() {
	if !auth() {
		fmt.Println("密码输入错误, 程序退出")
		return
	}
	operates := map[string]func(){
		"add":    addUser,
		"modify": modifyUser,
		"delete": deleteUser,
		"query":  queryUser,
	}
	for {
		text := input("请输入指令: ")
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
