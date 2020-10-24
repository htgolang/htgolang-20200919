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

func deleteUserById(id string) {
	// 方法1:
	// 把不需要删除 => 定义新的切片 => users
	// tempUsers := []map[string]string{}
	tempUsers := make([]map[string]string, 0, len(users)-1)
	for _, user := range users {
		if user["id"] != id {
			tempUsers = append(tempUsers, user)
		}
	}
	users = tempUsers

	// 方法2:
	// index => append([:n], [n+1:])
	// copy
}

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

func main() {
	fmt.Printf("%#v\n", users)
	deleteUser()
	fmt.Println(users)
}
