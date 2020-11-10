package funcs

import (
	"fmt"
	"homework/pkg/models"
)

func userModify() {
	var userID string
	fmt.Println("请输入要修改的用户ID")
	fmt.Scan(&userID)

	for k,v:=range models.Users {
		if v["ID"] == userID {
			fmt.Println("要修改的用户ID数据如下：")
			userIdList(k)
			//修改
			Name, Contact, Address := inputUser()
			models.Users[k]["Name"] = Name
			models.Users[k]["Contact"] = Contact
			models.Users[k]["Address"] = Address
			//打印修改后信息
			fmt.Println("用户ID数据变更如下：")
			userIdList(k)
			return
		}
	}
}