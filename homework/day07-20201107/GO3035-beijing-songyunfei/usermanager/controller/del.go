package controller

import (
	"fmt"
	"strconv"
	"usermanager/users"
	"usermanager/userutils"
)

func delUser(udb users.Mydb) {
	for {
		uid := userutils.Input("提示:按q 退出\n请输入用户ID:")
		if uid == "q"{
			break
		}
		fid,_ := strconv.Atoi(uid)
		uinfo,_,err := udb.FindByid(fid)
		if err != nil {
			fmt.Println(err)
		} else {
			header := []string{"ID","用户名","生日","联系电话","地址"}
			data := [][]string{{strconv.Itoa(uinfo.Id), uinfo.Name, uinfo.Birthday.Format("2006-1-2"),uinfo.Tel,uinfo.Addr},}
			userutils.Showintable(header,data)
			yOrn :=userutils.Input("确认删除?(q 退出.)[y/n/q]:")
			if yOrn == "y" {
				err = udb.Del(fid)
				if err != nil {
					fmt.Println(err)
				}
				err = udb.Sync()
				if err == nil {
					fmt.Println("已删除..")
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