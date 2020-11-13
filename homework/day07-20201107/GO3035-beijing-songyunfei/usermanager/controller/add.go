package controller

import (
	"fmt"
	"usermanager/users"
	"usermanager/userutils"
)

func Add(udb users.Mydb){

	for {
		u := userutils.Input("请输入用户名:")
		a := userutils.Input("请输入地址:")
		t := userutils.Input("请输入联系电话:")
		p := userutils.Input("请输入密码:")
		b := userutils.Input("请输入生日(示例:1999-01-05):")
		header := []string{"用户名","地址","联系电话","生日"}
		data := [][]string{{u,a,t,b}}
		userutils.Showintable(header,data)
		yorn := userutils.Input("确认添加y/n:")
		if yorn == "y" {
			if err := udb.Add(u,a,t,p,b); err != nil{
				fmt.Println(err)
			}else {
				if err = udb.Sync(); err == nil{
					fmt.Println("添加成功.")

				}

			}
			q := userutils.Input("q.退出, 任意键继续添加....")
			if q == "q"{
				break
			}else {
				continue
			}

		} else if yorn == "n" {
			q := userutils.Input("q.退出, 任意键继续....")
			if q == "q"{
				break
			}else {
				continue
			}
		}
	}
}
