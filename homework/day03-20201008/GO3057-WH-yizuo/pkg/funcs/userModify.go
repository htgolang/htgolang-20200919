package funcs

import (
	"GO3057-WH-yizuo/pkg/models"
	"fmt"
)

func UserModify() {
	var UserID string
	fmt.Println("请输入要变更的用户ID： ")
	fmt.Scan(&UserID)

	// 新增用户至数据
	for k, v := range models.Users {
		// 判断如果遍历的数据与需要删除的数据相同
		if v["ID"] == UserID {
			// 获取需要变更的用户数据
			Name, Contact, Address := InuptUsersElement()
			models.Users[k]["Name"] = Name
			models.Users[k]["Contact"] = Contact
			models.Users[k]["Address"] = Address
			fmt.Println("用户数据变更成功。")
		}
	}

}
