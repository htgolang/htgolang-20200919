package usertools

import (
	"fmt"
	"strconv"
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

// GetInfo is ..
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

// UserAdd is ...
func UserAdd() {
	uid := getUID()
	userinfo := make(map[string]string)

	username, userphone, useraddr := getInfo()

	userinfo["userid"] = strconv.Itoa(uid)
	userinfo["username"] = username
	userinfo["userphone"] = userphone
	userinfo["useraddr"] = useraddr

	usersInfo = append(usersInfo, userinfo)

	usertable := []string{strconv.Itoa(uid), username, userphone, useraddr}

	fmt.Println("\nEnter successfully! Your account information is: ")
	userTable(usertable)

	// fmt.Printf("输入成功，您的账号信息是：id是 %d, 用户名是 %s, 联系方式是 %s, 联系地址是 %s\n", uid, username, userphone, useraddr)

}
