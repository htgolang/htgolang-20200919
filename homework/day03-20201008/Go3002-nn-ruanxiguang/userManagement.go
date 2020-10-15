package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add() {
	fmt.Println("请按照如下内容填写信息")
	var (
		ID    int
		name  string
		phone string
		addr  string
	)
	fmt.Print("姓名：")
	fmt.Scan(&name)
	fmt.Print("\n联系方式(电话)：")
	fmt.Scan(&phone)
	fmt.Print("\n通信地址：")
	fmt.Scan(&addr)
	ID = len(users) + 1

	adduser := map[string]string{}
	adduser["ID"] = strconv.Itoa(ID)
	adduser["name"] = name
	adduser["phone"] = phone
	adduser["addr"] = addr
	fmt.Println("-----添加的用户信息-----")
	fmt.Printf("用户ID：%d\n", ID)
	fmt.Printf("名称：%s\n", name)
	fmt.Printf("联系方式：%s\n", phone)
	fmt.Printf("地址：%s\n", addr)
	fmt.Println("------------------------")
	var ok string
	fmt.Print("是否添加(Y/N):")
	fmt.Scan(&ok)
	if strings.ToLower(ok) == "y" {
		users = append(users, adduser)
		fmt.Printf("%s用户添加成功\n", name)
	} else {
		add()
	}
}

func del() {
	var ID string
	var OK string
	for _, v := range users {
		for id := range v {
			if id == "ID" {
				fmt.Printf("%s、%s\n", v[id], v["name"])
			}
		}
	}
	fmt.Print("输入要删除的用户ID：")
	fmt.Scan(&ID)
	for i, v := range users {
		if ID == v["ID"] {
			fmt.Printf("用户信息：%v, 是否删除(Y/y):", users[i])
			fmt.Scan(&OK)
			if strings.ToLower(OK) == "y" {
				if i == 0 {
					users = users[1:]
				} else if i == len(users)-1 {
					users = users[:len(users)-1]
				} else {
					users = append(users[:i], users[i+1:]...)
				}
			} else {
				fmt.Println("取消")
			}
		}
	}
	fmt.Println(users)
}

func modify() {
	var (
		ID    string
		name  string
		phone string
		addr  string
		OK    string
	)
	var flag bool
	for _, v := range users {
		for id := range v {
			if id == "ID" {
				fmt.Printf("%s、%s\n", v[id], v["name"])
			}
		}
	}
	fmt.Printf("修改用户ID：")
	fmt.Scan(&ID)
	for i, v := range users {
		if ID == v["ID"] {
			fmt.Printf("要修改的用户信息：%v, 是否修改(Y/y):", users[i])
			fmt.Scan(&OK)
			if strings.ToLower(OK) == "y" {
				fmt.Print("姓名：")
				fmt.Scan(&name)
				fmt.Print("\n联系方式(电话)：")
				fmt.Scan(&phone)
				fmt.Print("\n通信地址：")
				fmt.Scan(&addr)
				users[i]["name"] = name
				users[i]["phone"] = phone
				users[i]["addr"] = addr
			} else {
				fmt.Println("取消")
			}
		}
	}
	if flag == false {
		fmt.Println("ID不存在")
	}
	fmt.Println(users)

}

func query() {
	var info string
	var flag bool
	fmt.Print("请输入要查找的信息：")
	fmt.Scan(&info)
	for i, v := range users {
		flag = false
		for _, in := range v {
			if strings.Contains(in, info) && flag == false {
				flag = true
				fmt.Println(users[i])
			}
		}
	}

}

var users = []map[string]string{}

func main() {
	for {
		fmt.Println(`
----------用户管理---------
		1、添加用户
		2、删除用户
		3、修改用户
		4、查找用户
		5、退出(q键)
---------------------------`)
		fmt.Printf("%v", users)
		var userchoice string
		fmt.Print("请选择：")
		fmt.Scan(&userchoice)
		switch {
		case userchoice == "1":
			add()
		case userchoice == "2":
			del()
		case userchoice == "3":
			modify()
		case userchoice == "4":
			query()
		case strings.ToLower(userchoice) == "q":
			os.Exit(0)
		default:
			fmt.Println("请选择(1/2/3/4)")
		}
	}

}
