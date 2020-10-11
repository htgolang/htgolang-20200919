package main

import (
	"fmt"
	"os"
	"strconv"
)

func add() {
	fmt.Println("add")
	var (
		ID    string
		name  string
		phone string
		addr  string
		ok    string
	)
	fmt.Printf("姓名：")
	fmt.Scan(&name)
	fmt.Printf("\n电话：")
	fmt.Scan(&phone)
	fmt.Printf("\n地址：")
	fmt.Scan(&addr)
	//定义ID
	if len(users) == 0 {
		ID = strconv.Itoa(1)
	} else {
		var tmpID int
		for _, v := range users {
			if res, err := strconv.Atoi(v["ID"]); err == nil {
				if tmpID < res {
					tmpID = res
				}
			} else {
				fmt.Println("用户中有ID值错误")
			}
		}

		ID = strconv.Itoa(tmpID + 1)
	}

	adduser := map[string]string{}
	adduser["ID"] = ID
	adduser["name"] = name
	adduser["phone"] = phone
	adduser["addr"] = addr
	fmt.Printf("添加的信息为： %v\n", adduser)
	fmt.Printf("确认是否添加, (y/Y)添加")
	fmt.Scan(&ok)
	if ok == "y" || ok == "Y" {
		users = append(users, adduser)
	} else {
		add()
	}
}
func del() {
	var ID string
	var ok string
	fmt.Printf("输入要删除的用户ID: ")
	fmt.Scan(&ID)

	for i, v := range users {
		if v["ID"] == ID {
			fmt.Printf("用户信息：%v, 是否确认删除(y/Y): ", users[i])
			fmt.Scan(&ok)
			if ok == "y" || ok == "Y" {
				if i == 0 {
					users = users[1:]
				} else if i == len(users)-1 {
					users = users[:len(users)-1]
				} else {
					users = append(users[:i], users[i+1:]...)
				}
			} else {
				fmt.Println("删除取消。")
			}
		}
	}
	fmt.Println(users)

}

func modify() {
	fmt.Println("modify")
	var ID, name, phone, addr, ok string
	var flage bool
	fmt.Printf("要修改的用户id：")
	fmt.Scan(&ID)
	for i, v := range users {
		if v["ID"] == ID {
			flage = true
			fmt.Printf("要修改的用户信息 %v， 是否修改(Y or y/N or n)\n", users[i])
			fmt.Scan(&ok)
			if ok == "Y" || ok == "y" {
				fmt.Printf("姓名：")
				fmt.Scan(&name)
				fmt.Printf("\n电话：")
				fmt.Scan(&phone)
				fmt.Printf("\n地址：")
				fmt.Scan(&addr)
				users[i]["name"] = name
				users[i]["phone"] = phone
				users[i]["addr"] = addr
			} else {
				fmt.Println("修改撤销。")
			}
		}
	}
	if flage == false {
		fmt.Println("ID不存在。")
	}
	fmt.Println(users)
}

func query() {
	fmt.Println("query")
	var info string
	fmt.Printf("输入要查找的信息：")
	fmt.Scan(&info)
	tmpuser := []map[string]string{}
	for i, v := range users {
		for _, in := range v {
			if in == info {
				tmpuser = append(tmpuser, users[i])
			}
		}
	}
	fmt.Println(tmpuser)
}

var users = []map[string]string{}

func main() {
	for {
		fmt.Printf(`
		--------------------
		|  1、添加用户      |
		|  2、删除用户      |
		|  3、修改用户      |
		|  4、查找用户      |
		|  q、退出(exit)   |
		--------------------`)
		fmt.Println("")
		var userchoice string
		fmt.Printf("输入你的选择: ")
		fmt.Scan(&userchoice)
		switch {
		case userchoice == "1":
			add()
			fmt.Println(users)
		case userchoice == "2":
			del()
		case userchoice == "3":
			modify()
		case userchoice == "4":
			query()
		case userchoice == "q" || userchoice == "exit":
			os.Exit(0)
		default:
			fmt.Println("输入(1/2/3/4选项！！！)")
		}
	}
}
