package usertools

import (
	"fmt"
	"strconv"
)

// UserDel is ..
func UserDel() {
	delID, _ := fmt.Print("请输入要删除的用户id: ")
	fmt.Scan(&delID)

	for i, userinfo := range usersInfo {
		if strconv.Itoa(delID) == userinfo["userid"] {
			usertable := []string{userinfo["userid"], userinfo["username"], userinfo["userphone"], userinfo["useraddr"]}
			userTable(usertable)
			// fmt.Printf("用户信息为用户 %s 联系方式是 %s 联系地址是 %s \n", userinfo["username"], userinfo["userphone"], userinfo["useraddr"])
			var judge string
			fmt.Print("是否该删除该用户，y/n: ")
			fmt.Scan(&judge)
			for {
				if judge == "y" {
					usersInfo = append(usersInfo[:i], usersInfo[i+1:]...)
					fmt.Println("Delete successfully!")
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
