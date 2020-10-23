package main

import (
	"fmt"
	"strconv"
	"strings"
)

func printMsg() {
	fmt.Printf(`
	--------------------
	  1、查看用户      
	  2、添加用户      
	  3、删除用户      
	  4、修改用户      
	  5、查找用户      
	  h、查看帮助      
	  q、退出(exit)   
	--------------------
	`)
}

func printUsers(users []map[string]string) {
	//打印用户信息
	fmt.Println("\n以下是目前已存在的用户：")
	for _, v := range users {

		fmt.Printf("ID：%s, 名称：%s, 联系方式：%s, 通信地址：%s\n", v["ID"], v["name"], v["phone"], v["address"])

	}
}

func addUser(users []map[string]string) {
	var (
		name    string
		phone   string
		address string
		ID      string
	)
	fmt.Println("请按顺序和格式输入你要增加的用户信息（名称 联系方式 通信地址），例如：zhangsan 1885642880 Guangzhou")
	fmt.Scanln(&name, &phone, &address)

	// maxIDtmp1 := users[len(users)-1]["ID"]
	maxIDtmp2, err := strconv.Atoi(users[len(users)-1]["ID"])
	if err != nil {
		panic(err)
	}
	ID = strconv.Itoa(maxIDtmp2 + 1)
	users = append(users, map[string]string{"ID": ID, "name": name, "phone": phone, "address": address})
	printUsers(users)

}

func delUser(users []map[string]string) {
	var (
		ID string
		yn string
	)
	fmt.Println("请输入要删除的用户 ID：")
	fmt.Scanln(&ID)

	for k, v := range users {
		st, ok := v["ID"]
		// fmt.Println(st)
		if ok && st == ID {
			fmt.Println("该用户存在！！！")
			fmt.Printf("ID：%s, 名称：%s, 联系方式：%s, 通信地址：%s\n", v["ID"], v["name"], v["phone"], v["address"])
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
		ID      string
		yn      string
		name    string
		phone   string
		address string
	)
	fmt.Println("请输入要删除的用户 ID：")
	fmt.Scanln(&ID)

	for _, v := range users {
		st, ok := v["ID"]
		// fmt.Println(st)
		if ok && st == ID {
			fmt.Println("该用户存在！！！")
			fmt.Printf("ID：%s, 名称：%s, 联系方式：%s, 通信地址：%s\n", v["ID"], v["name"], v["phone"], v["address"])
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
				fmt.Scanln(&phone)
				if len(phone) != 0 {
					v["phone"] = phone
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
				fmt.Printf("ID：%s, 名称：%s\n", v["ID"], v["name"])
			}
		}
	}
	// printUsers(users)
}

func main() {
	users := []map[string]string{
		{"ID": "1", "name": "test1", "phone": "11234", "address": "Guangdong"},
		{"ID": "2", "name": "test2", "phone": "21234", "address": "Shanghai"},
		{"ID": "3", "name": "test3", "phone": "31234", "address": "Hangzhou"},
		{"ID": "4", "name": "test4", "phone": "41234", "address": "Shenzhen"},
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
		//fallthrough
	case "q", "quit":
		break
	default:
		fmt.Println("操作错误")
	}
}
