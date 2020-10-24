package main

import (
	"fmt"
	"strings"
)

var users = []map[string]string{}

func init() {
	for i := 0; i < 3; i++ {
		users = append(users, map[string]string{
			"id":   fmt.Sprintf("%d", i),
			"name": fmt.Sprintf("name_%d", i),
			"addr": fmt.Sprintf("addr_%d", i),
			"tel":  fmt.Sprintf("tel_%d", i),
		})
	}
}

func input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return text
}

func findUserById(id string) map[string]string {
	for _, user := range users {
		if user["id"] == id {
			return user
		}
	}
	return nil
}

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

func main() {
	fmt.Printf("%#v\n", users)
	modifyUser()
	fmt.Println(users)
}
