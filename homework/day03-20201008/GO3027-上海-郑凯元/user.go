package main

import (
	"fmt"
	"strconv"
	"strings"
)

var usersInfo = []map[string]string{
	{"userid": "1", "username": "zky1", "userphone": "1234567891", "useraddr": "sh1"},
	{"userid": "2", "username": "zky2", "userphone": "1234567892", "useraddr": "sh2"},
	{"userid": "3", "username": "zky3", "userphone": "1234567893", "useraddr": "sh3"},
}

func getUID() int {
	uid := 1

	if len(usersInfo) == 0 {

	} else {
		for _, v := range usersInfo {
			if uuid, _ := strconv.Atoi(v["userid"]); uuid >= uid {
				uid = uuid + 1
			}
		}

	}
	return uid
}

func getInfo() (string, string, string) {
	var username, userphone, useraddr string

	fmt.Print("请输入你的用户名：")
	fmt.Scan(&username)

LABEL1:
	for _, userinfo := range usersInfo {
		if userinfo["username"] == username {
			fmt.Print("用户名重复，请重新输入你的用户名：")
			fmt.Scan(&username)
			goto LABEL1
		}
	}

	fmt.Print("请输入你的联系方式：")
	fmt.Scan(&userphone)

	for {
		if _, err := strconv.Atoi(userphone); err != nil {
			fmt.Print("请输入正确联系方式：")
			fmt.Scan(&userphone)
		} else {
			break
		}
	}

	fmt.Print("请输入你的联系地址：")
	fmt.Scan(&useraddr)
	return username, userphone, useraddr
}

func userAdd() {
	uid := getUID()
	userinfo := make(map[string]string)

	username, userphone, useraddr := getInfo()

	userinfo["userid"] = strconv.Itoa(uid)
	userinfo["username"] = username
	userinfo["userphone"] = userphone
	userinfo["useraddr"] = useraddr

	usersInfo = append(usersInfo, userinfo)

	fmt.Printf("输入成功，您的账号信息是：id是 %d, 用户名是 %s, 联系方式是 %s, 联系地址是 %s\n", uid, username, userphone, useraddr)

}

func userDel() {
	delID, _ := fmt.Print("请输入要删除的用户id: ")
	fmt.Scan(&delID)

	for i, userinfo := range usersInfo {
		if strconv.Itoa(delID) == userinfo["userid"] {
			fmt.Printf("用户信息为用户 %s 联系方式是 %s 联系地址是 %s \n", userinfo["username"], userinfo["userphone"], userinfo["useraddr"])
			var judge string
			fmt.Print("是否该删除该用户，y/n: ")
			fmt.Scan(&judge)
			for {
				if judge == "y" {
					usersInfo = append(usersInfo[:i], usersInfo[i+1:]...)
					fmt.Println(usersInfo)
					break
				} else if judge == "n" {
					break
				} else {
					fmt.Print("请输入正确的字母，y/n: ")
					fmt.Scan(&judge)
				}
			}
		}
	}
}

func userModify() {
	modifyID, _ := fmt.Print("请输入要修改的用户id: ")
	fmt.Scan(&modifyID)

	for _, userinfo := range usersInfo {
		if strconv.Itoa(modifyID) == userinfo["userid"] {
			fmt.Printf("用户信息为用户 %s 联系方式是 %s 联系地址是 %s \n", userinfo["username"], userinfo["userphone"], userinfo["useraddr"])
			var modify string
			fmt.Print("是否该修改该用户，y/n: ")
			fmt.Scan(&modify)
			for {
				if modify == "y" {
					username, userphone, useraddr := getInfo()
					userinfo["username"] = username
					userinfo["userphone"] = userphone
					userinfo["useraddr"] = useraddr
					fmt.Println(usersInfo)
					break
				} else if modify == "n" {
					break
				} else {
					fmt.Print("请输入正确的字母，y/n: ")
					fmt.Scan(&modify)
				}
			}
		}
	}
}

func userQuery() {
	var queryinfo string
	fmt.Print("请输入要查询的内容：")
	fmt.Scan(&queryinfo)

	s := []string{"username", "userphone", "useraddr"}

	for userindex, userinfo := range usersInfo {
		for _, ss := range s {
			if strings.Contains(userinfo[ss], queryinfo) {
				fmt.Println(usersInfo[userindex])
			}
		}
	}
}

func main() {
	// userAdd()
	// userDel()
	// userModify()
	// userQuery()
	i := true

	for i {
		var choose string
		fmt.Print("1: 添加用户\n2: 删除用户\n3: 修改用户\n4: 查询用户\n5：退出\n请输入你要做的事情：")
		fmt.Scan(&choose)

		switch {
		case choose == "1":
			userAdd()
		case choose == "2":
			userDel()
		case choose == "3":
			userModify()
		case choose == "4":
			userQuery()
		case choose == "5":
			i = false
		}
	}

}
