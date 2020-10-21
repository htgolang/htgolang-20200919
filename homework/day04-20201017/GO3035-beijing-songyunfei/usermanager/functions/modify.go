package functions

import (
	"fmt"
	"strconv"
	"usermanager/users"
)

// 修改用户函数
func ModifyUser()  {
	for {
		var input string
		fmt.Printf("提示:按q 退出\n请输入用户ID:")
		_, _ = fmt.Scanln(&input)
		if input == "q" {
			break
		}
		uid, _ := strconv.Atoi(input)
		index, err := findByid(uid)
		if err != nil {
			fmt.Println("未找到ID:",uid)
			continue
		}
		header := []string{"ID","用户名","联系电话","地址"}
		data := [][]string{{users.Users[index]["id"], users.Users[index]["name"], users.Users[index]["tel"], users.Users[index]["addr"]},}
		showintable(header,data)
		var name,tel,addr,yOrn string
		fmt.Printf("添加用户:\n请输入新用户名:")
		_, _ = fmt.Scanln(&name)
		fmt.Printf("请输入新联系电话:")
		_, _ = fmt.Scanln(&tel)
		fmt.Printf("请输入新地址:")
		_, _ = fmt.Scanln(&addr)
		nheader := []string{"用户名","联系电话","地址"}
		ndata := [][]string{{fmt.Sprintf("%s ---> %s",users.Users[index]["name"], name),
			fmt.Sprintf("%s ---> %s",users.Users[index]["tel"], tel),
			fmt.Sprintf("%s ---> %s",users.Users[index]["addr"], addr)},}
		showintable(nheader,ndata)
		fmt.Printf("确认以上修改?[y/n]:")
		_,_ = fmt.Scanln(&yOrn)
		if yOrn == "y"{
			users.Users[index]["name"] = name
			users.Users[index]["tel"] = tel
			users.Users[index]["addr"] = addr
			fmt.Printf("修改成功. 继续请按 c. 任意键退出...\n")
			var t string
			_,_ = fmt.Scanln(&t)
			if t == "c"{
				continue
			}
			break

		}else if yOrn == "n"{
			continue
		}
	}
}
