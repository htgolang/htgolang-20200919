package functions

import (
	"fmt"
	"strconv"
	"usermanager/users"
)


//添加用户函数
func Useradd() {
	for {
		var name, tel, addr, yesOrno string
		fmt.Printf("添加用户:\n请输入用户名:")
		_, _ = fmt.Scanln(&name)
		fmt.Printf("请输入联系电话:")
		_, _ = fmt.Scanln(&tel)
		fmt.Printf("请输入地址:")
		_, _ = fmt.Scanln(&addr)
		header := []string{"用户名","联系电话","地址"}
		data := [][]string{{name, tel, addr},}
		showintable(header,data)
		fmt.Printf("确认添加?[y/n]:")
		_, _ = fmt.Scanln(&yesOrno)
		if yesOrno == "y" {
			// 调用添加用户的函数
			//users = addUser(name, tel, addr, users)
			aduser := make(map[string]string)
			id := func() int {
				if len(users.Users) == 0 {
					return 1
				}
				id := 1
				for i := 0; i < len(users.Users); i++{
					tid,_ := strconv.Atoi(users.Users[i]["id"])
					if id < tid{
						id = tid
					}
				}
				return id +1
			}()
			aduser["id"] = strconv.Itoa(id)
			aduser["name"] = name
			aduser["tel"] = tel
			aduser["addr"] = addr
			users.Users = append(users.Users,aduser)
			var q string
			fmt.Printf("q.退出, 任意键继续....")
			_, _ = fmt.Scanln(&q)
			if q == "q"{
				break
			}else {
				continue
			}

		} else if yesOrno == "n" {
			var q string
			fmt.Printf("q.退出, 任意键继续....")
			_, _ = fmt.Scanln(&q)
			if q == "q"{
				break
			}else {
				continue
			}
		}
	}
}
