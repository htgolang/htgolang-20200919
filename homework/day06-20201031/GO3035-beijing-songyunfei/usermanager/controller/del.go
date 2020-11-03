package controller

import (
	"fmt"
	"strconv"
	"usermanager/users"
	"usermanager/userutils"
)

func delUser(udb *users.Userdb) {
	for {
		uid := userutils.Input("提示:按q 退出\n请输入用户ID:")
		if uid == "q"{
			break
		}
		fid,_ := strconv.Atoi(uid)
		index,err := udb.FindByid(fid)
		if err != nil {
			fmt.Println(err)
		} else {
			header := []string{"ID","用户名","生日","联系电话","地址"}
			data := [][]string{{strconv.Itoa(udb.UserSlice[index].Id), udb.UserSlice[index].Name, udb.UserSlice[index].Birthday.Format("2006-1-2"),udb.UserSlice[index].Tel,udb.UserSlice[index].Addr},}
			userutils.Showintable(header,data)
			yOrn :=userutils.Input("确认删除?(q 退出.)[y/n/q]:")
			if yOrn == "y" {
				err = udb.Del(fid)
				if err != nil {
					fmt.Println(err)
				}
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