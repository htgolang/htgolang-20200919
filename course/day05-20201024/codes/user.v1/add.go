package main

import (
	"fmt"
	"strconv"
)

// 定义存储用户的切片
// 每个元素为一个用户
// 用户包含多个属性, 用户使用映射进行存储
// key => string, value => string
// id => 编号
// name => 名称
// addr => 地址
// tel => 电话号码

var users = []map[string]string{}

// 获取增加用户的ID值
func getId() string {
	// []int 值最大的元素
	// [5, 2, 9, 3, 6]
	id := 0
	for _, user := range users {
		if uid, err := strconv.Atoi(user["id"]); err == nil {
			if uid > id {
				id = uid // 当找到比当前记录值大, 更新最大值
			}
		}
	}
	return strconv.Itoa(id + 1)
}

func input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return text
}

func addUser() {
	user := map[string]string{
		"id":   getId(),
		"name": input("请输入用户名: "),
		"addr": input("请输入联系地址: "),
		"tel":  input("请输入联系方式: "),
	}

	users = append(users, user)
}

func main() {
	addUser()
	fmt.Printf("%#v\n", users)
}
