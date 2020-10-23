package functions

import (
	"fmt"
	"strconv"
	"usermanager/users"
)

// 删除用户函数
func DelUser() {
	for {
		var uid string
		fmt.Printf("提示:按q 退出\n请输入用户ID:")
		_, _ = fmt.Scanln(&uid)
		if uid == "q"{
			break
		}
		fid,_ := strconv.Atoi(uid)
		index, err := findByid(fid)
		fmt.Println(index)
		if err != nil {
			fmt.Println("未找到用户id:", uid)
		} else {
			//fmt.Printf("Id: %s, 用户名: %s, 联系方式: %s, 地址:%s.\n确认删除?(q 退出.)[y/n/q]", users.Users[index]["id"], users.Users[index]["name"], users.Users[index]["tel"], users.Users[index]["addr"])
			header := []string{"ID","用户名","联系电话","地址"}
			data := [][]string{{users.Users[index]["id"], users.Users[index]["name"], users.Users[index]["tel"], users.Users[index]["addr"]},}
			showintable(header,data)
			fmt.Printf("确认删除?(q 退出.)[y/n/q]:")
			var yOrn string
			_, _ = fmt.Scanln(&yOrn)
			if yOrn == "y" {
				if index == 0 && len(users.Users) == 1{
					users.Users = []map[string]string{}
					break
				}
				if index == len(users.Users) -1 {
					users.Users = users.Users[:index]
					break
				}
				users.Users = append(users.Users[:index], users.Users[index+1:]...)
				break
			}else if yOrn == "n"{
				continue
			}else if yOrn == "q"{
				break
			}else {
				break
			}
		}
	}
}