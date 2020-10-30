package controller

import (
	"fmt"
	"usermanager/users"
	"usermanager/userutils"
)

func queryuser(udb *users.Userdb) {
	str := userutils.Input("请出入关键字:")
	u,ok := udb.QueryUser(str)
	if ok {
		header := []string{"用户名","地址","联系电话","生日"}
		data := [][]string{{u.Name,u.Addr,u.Tel,u.Birthday.Format("2006-1-2")}}
		userutils.Showintable(header,data)
	}else {
		fmt.Println("未找到.....")
	}
}
