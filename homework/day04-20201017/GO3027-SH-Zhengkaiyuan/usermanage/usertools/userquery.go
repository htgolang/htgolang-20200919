package usertools

import (
	"fmt"
	"strings"
)

// UserQuery is ..
func UserQuery() {
	var queryinfo string
	fmt.Print("请输入要查询的内容：")
	fmt.Scan(&queryinfo)

	s := []string{"username", "userphone", "useraddr"}

	for _, userinfo := range usersInfo {
		for _, ss := range s {
			if strings.Contains(userinfo[ss], queryinfo) {
				usertable := []string{userinfo["userid"], userinfo["username"], userinfo["userphone"], userinfo["useraddr"]}
				userTable(usertable)
			}

		}
	}
}
