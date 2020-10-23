package usertools

import (
	"fmt"
	"strconv"
)

// UserModify is ..
func UserModify() {
	modifyID, _ := fmt.Print("请输入要修改的用户id: ")
	fmt.Scan(&modifyID)

	for _, userinfo := range usersInfo {
		if strconv.Itoa(modifyID) == userinfo["userid"] {
			usertable := []string{userinfo["userid"], userinfo["username"], userinfo["userphone"], userinfo["useraddr"]}
			userTable(usertable)

			var modify string
			fmt.Print("是否该修改该用户，y/n: ")
			fmt.Scan(&modify)
			for {
				if modify == "y" {
					username, userphone, useraddr := getInfo()
					userinfo["username"] = username
					userinfo["userphone"] = userphone
					userinfo["useraddr"] = useraddr
					fmt.Println("\nModified successfully! New user information is：")
					usertable := []string{userinfo["userid"], username, userphone, useraddr}
					userTable(usertable)
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
