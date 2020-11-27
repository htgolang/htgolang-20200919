package controller

import (
	"fmt"
	"strconv"
	"usermanager_net/users"
	"usermanager_net/userutils"
)

func modifyUser(udb users.Mydb)  {
	for {
		input := userutils.Input("提示:按q 退出\n请输入用户ID:")
		if input == "q" {
			break
		}
		uid, _ := strconv.Atoi(input)
		uinfo, index, err := udb.FindByid(uid)
		if err != nil {
			fmt.Println(err)
			continue
		}
		header := []string{"ID","用户名","生日","联系电话","地址"}
		data := [][]string{{strconv.Itoa(uinfo.Id), uinfo.Name, uinfo.Birthday.Format("2006-1-2"),uinfo.Tel,uinfo.Addr},}
		userutils.Showintable(header,data)
		name := userutils.Input("添加用户:\n请输入新用户名:")
		addr := userutils.Input("请输入地址:")
		tel := userutils.Input("请输入联系电话:")
		//p := userutils.Input("请输入密码:")
		b := userutils.Input("请输入生日(示例:1999-01-05):")

		nheader := []string{"用户名","生日","联系电话","地址"}
		ndata := [][]string{{fmt.Sprintf("%s ---> %s",uinfo.Name, name),
			fmt.Sprintf("%s ---> %s",uinfo.Tel, tel),
			fmt.Sprintf("%s ---> %s",uinfo.Addr, addr),
			fmt.Sprintf("%s ---> %s",uinfo.Birthday.Format("2006-1-2"),b)},}
		userutils.Showintable(nheader,ndata)
		yOrn := userutils.Input("确认以上修改?[y/n]:")
		if yOrn == "y"{
			err := udb.Modify(index,name,addr,tel,b)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if err = udb.Sync(); err != nil {
				fmt.Println("同步失败",err)
			}
			t := userutils.Input("修改成功. 继续请按 c. 任意键退出...\n")
			if t == "c"{
				continue
			}
			break

		}else if yOrn == "n"{
			continue
		}
	}
}
