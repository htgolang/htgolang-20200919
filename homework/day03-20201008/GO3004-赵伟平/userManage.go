package main

import (
	"fmt"
	"strconv"
	"strings"
)

func printMsg() {
	fmt.Println("============================")
	fmt.Println("=========== 用户管理系统  ===========")
	fmt.Println("==      请根据以下提示信息操作       ==")

	fmt.Println("==== 查看用户：输入 1 或者 all 后按回车 ====")

	fmt.Println("==== 添加用户：输入 2 或者 add 后按回车 ====")

	fmt.Println("==== 删除用户：输入 3 或者 del 后按回车 ====")

	fmt.Println("==== 修改用户：输入 4 或者 modify 后按回车 ====")

	fmt.Println("==== 查找用户：输入 5 或者 query 后按回车 ====")

	fmt.Println("==== 查看帮助：输入 h 或者 help 后按回车  ==")

	fmt.Println("==== 退出系统：输入 q 或者 quit 后按回车  ==")
}

func printUsers(users []map[string]string) {
	//打印用户信息
	fmt.Println("\n以下是目前已存在的用户：")
	for _, v := range users {

		fmt.Printf("ID：%s, 名称：%s, 联系方式：%s, 通信地址：%s\n", v["id"], v["name"], v["contact"], v["address"])

	}
}

func addUser(users []map[string]string) {
	var (
		name    string
		contact string
		address string
		id      string
	)
	fmt.Println("请按顺序和格式输入你要增加的用户信息（名称 联系方式 通信地址），例如：Aomine 123123 Guangzhou")
	fmt.Scanln(&name, &contact, &address)

	// maxIDtmp1 := users[len(users)-1]["id"]
	maxIDtmp2, err := strconv.Atoi(users[len(users)-1]["id"])
	if err != nil {
		panic(err)
	}
	id = strconv.Itoa(maxIDtmp2 + 1)
	users = append(users, map[string]string{"id": id, "name": name, "contact": contact, "address": address})
	printUsers(users)

}

func delUser(users []map[string]string) {
	var (
		id string
		yn string
	)
	fmt.Println("请输入要删除的用户 id：")
	fmt.Scanln(&id)

	for k, v := range users {
		st, ok := v["id"]
		// fmt.Println(st)
		if ok && st == id {
			fmt.Println("该用户存在！！！")
			fmt.Printf("ID：%s, 名称：%s, 联系方式：%s, 通信地址：%s\n", v["id"], v["name"], v["contact"], v["address"])
			fmt.Println("请输入 y/n 确认是否删除：")
			fmt.Scanln(&yn)
			if yn == "y" {
				users = append(users[:k], users[k+1:]...)
			} else if yn == "n" {
				continue
			} else {
				fmt.Println("输入错误，请重新开始")
			}
		} else {
			fmt.Println("该用户不存在！！！")
			break
		}
	}
	printUsers(users)
}

func modifyUser(users []map[string]string) {
	var (
		id      string
		yn      string
		name    string
		contact string
		address string
	)
	fmt.Println("请输入要删除的用户 id：")
	fmt.Scanln(&id)

	for _, v := range users {
		st, ok := v["id"]
		// fmt.Println(st)
		if ok && st == id {
			fmt.Println("该用户存在！！！")
			fmt.Printf("ID：%s, 名称：%s, 联系方式：%s, 通信地址：%s\n", v["id"], v["name"], v["contact"], v["address"])
			fmt.Println("请输入 y/n 确认是否修改：")
			fmt.Scanln(&yn)
			if yn == "y" {
				fmt.Println("请按顺序和格式输入你要修改的用户信息（名称 联系方式 通信地址），如果为空直接回车默认不修改：")
				fmt.Println("名称：")
				fmt.Scanln(&name)
				// fmt.Println(name)
				// fmt.Println(len(name))
				if len(name) != 0 {
					v["name"] = name
				}
				fmt.Println("联系方式：")
				fmt.Scanln(&contact)
				if len(contact) != 0 {
					v["contact"] = contact
				}
				fmt.Println("通信地址：")
				fmt.Scanln(&address)
				if len(address) != 0 {
					v["address"] = address
				}

			} else if yn == "n" {
				continue
			} else {
				fmt.Println("输入错误，请重新开始")
			}
		} else {
			fmt.Println("该用户不存在！！！")
			break
		}
	}
	printUsers(users)
}

func queryUser(users []map[string]string) {
	var queryS string
	fmt.Println("请输入你要查找的关键字：")
	fmt.Scanln(&queryS)
	// fmt.Println(queryS)
	for _, v := range users {
		for _, i := range v {
			// fmt.Println(k, i)
			if find := strings.Contains(i, queryS); find {
				fmt.Printf("ID：%s, 名称：%s\n", v["id"], v["name"])
			}
		}
	}
	// printUsers(users)
}

func main() {
	users := []map[string]string{
		{"id": "1", "name": "test1", "contact": "11234", "address": "Guangdong"},
		{"id": "2", "name": "test2", "contact": "21234", "address": "Shanghai"},
		{"id": "3", "name": "test3", "contact": "31234", "address": "Hangzhou"},
		{"id": "4", "name": "test4", "contact": "41234", "address": "Shenzhen"},
	}

	var operate string
	printMsg()
	fmt.Scanln(&operate)
	switch operate {
	case "1", "all":
		printUsers(users)
	case "2", "add":
		addUser(users)
	case "3", "del":
		delUser(users)
	case "4", "modify":
		modifyUser(users)
	case "5", "query":
		queryUser(users)
	case "h", "help":
		printMsg()
		fallthrough
	case "q", "quit":
		break
	default:
		fmt.Println("操作错误")
	}
}
