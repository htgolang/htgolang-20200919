package main

import (
	"fmt"
	"strings"
)

var users = []map[string]string{}

func init() {
	for i := 0; i < 20; i++ {
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

func filter(user map[string]string, q string) bool {
	return strings.Contains(user["name"], q) ||
		strings.Contains(user["addr"], q) ||
		strings.Contains(user["tel"], q)
}

func printUser(user map[string]string) {
	fmt.Println(user["id"])
}

func query() {
	q := input("请输入查询信息: ")
	fmt.Println("查询结果")
	for _, user := range users {
		if filter(user, q) {
			printUser(user)
		}
	}
}

func main() {
	fmt.Printf("%#v\n", users)
	query()
}
