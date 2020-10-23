package functions

import (
	"fmt"
	"strings"
	"usermanager/users"
)

//关键字查找用户函数
func QueryUser()  {
	var str string
	var f int
	fmt.Printf("请出入关键字:")
	_,_ = fmt.Scanln(&str)
	for _,v := range users.Users {
		if strings.Contains(v["name"],str) || strings.Contains(v["tel"],str)|| strings.Contains(v["addr"],str) {
			fmt.Println("找到了!!!")
			header := []string{"ID","用户名","联系电话","地址"}
			data := [][]string{{v["id"],v["name"],v["tel"],v["addr"]},}
			showintable(header,data)
			f = 1
			break
		}
	}
	if f == 0 {
		fmt.Println("未找到~~~~~~~~~~~~~~~~~~~")
	}

}
