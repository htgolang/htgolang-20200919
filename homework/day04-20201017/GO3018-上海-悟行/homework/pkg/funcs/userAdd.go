package funcs

import (
	"fmt"
	"homework/pkg/models"
)

func userAdd() {
	//获取用户数据，并转换为字典，添加到数据中
	ID:=findMaxID()
	Name,Contact,Address := inputUser()
	models.AppendElement(ID,Name,Contact,Address)
	fmt.Printf("用户%s添加成功！现用户数据如下\n",Name)
	//显示当前数据
	usersList(&models.Users)
}