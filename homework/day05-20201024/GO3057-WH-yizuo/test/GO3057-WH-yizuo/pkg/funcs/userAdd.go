package funcs

import (
	"GO3057-WH-yizuo/pkg/models"
	"fmt"
)

func UserAdd() {
	/*
		获取用户数据的值，并转换为字典，添加到我们的数据中
	*/
	// 获取最新可以使用的ID值(FindLargestElementID 会查找所有数据中的ID中最大的值+1返回)
	ID := FindLargestElementID()
	// 获取新增的用户数据
	Name, Contact, Address := InuptUsersElement()
	// 新增用户至数据
	models.AppendElement(ID, Name, Contact, Address)
	fmt.Printf("用户%s创建成功！现用户数据如下\n", Name)

	// 查看当前有多少数据
	UsersList(&models.Users)
}
