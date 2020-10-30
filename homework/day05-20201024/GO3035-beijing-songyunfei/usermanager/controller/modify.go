package controller

import (
	"fmt"
	"strconv"
	"usermanager/users"
	"usermanager/userutils"
)

func modifyUser(udb *users.Userdb)  {
	for {
		input := userutils.Input("提示:按q 退出\n请输入用户ID:")
		if input == "q" {
			break
		}
		uid, _ := strconv.Atoi(input)
		index, err := udb.FindByid(uid)
		if err != nil {
			fmt.Println(err)
			continue
		}
		header := []string{"ID","用户名","生日","联系电话","地址"}
		data := [][]string{{strconv.Itoa(udb.UserSlice[index].Id), udb.UserSlice[index].Name, udb.UserSlice[index].Birthday.Format("2006-1-2"),udb.UserSlice[index].Tel,udb.UserSlice[index].Addr},}
		userutils.Showintable(header,data)
		name := userutils.Input("添加用户:\n请输入新用户名:")
		addr := userutils.Input("请输入地址:")
		tel := userutils.Input("请输入联系电话:")
		//p := userutils.Input("请输入密码:")
		b := userutils.Input("请输入生日(示例:1999-01-05):")

		nheader := []string{"用户名","生日","联系电话","地址"}
		ndata := [][]string{{fmt.Sprintf("%s ---> %s",udb.UserSlice[index].Name, name),
			fmt.Sprintf("%s ---> %s",udb.UserSlice[index].Tel, tel),
			fmt.Sprintf("%s ---> %s",udb.UserSlice[index].Addr, addr),
			fmt.Sprintf("%s ---> %s",udb.UserSlice[index].Birthday.Format("2006-1-2"),b)},}
		userutils.Showintable(nheader,ndata)
		yOrn := userutils.Input("确认以上修改?[y/n]:")
		if yOrn == "y"{
			err := udb.Modify(index,name,addr,tel,b)
			if err != nil {
				fmt.Println(err)
				continue
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
