package funcs

import (
	"GO3057-WH-yizuo/pkg/models"
	"fmt"
)

func UserModify() {
	/*
		用户输入用户ID，根据用户ID变更对应的条目数据
	*/
	// 获取用户输入的用户ID
	var UserID string
	fmt.Println("请输入要变更的用户ID： ")
	fmt.Scan(&UserID)

	// 新增用户至数据
	for k, v := range models.Users {
		// 判断如果遍历的数据与需要删除的数据相同
		if v["ID"] == UserID {
			// 打印要变更的内容
			fmt.Println("您要变更的数据如下:")
			userIdList(k)

			// 获取需要变更的用户数据并修改
			Name, Contact, Address := InuptUsersElement()
			models.Users[k]["Name"] = Name
			models.Users[k]["Contact"] = Contact
			models.Users[k]["Address"] = Address

			// 打印变更后的数据
			fmt.Println("用户数据变更如下:。")
			userIdList(k)

			return
		}
	}

}
