package funcs

import (
	"GO3057-WH-yizuo/pkg/models"
	"fmt"
)

func UserQuery() {
	/*
		根据ID或者用户名或者电话或者邮箱地址来查找对应的用户信息
		获取用户输入的值来判断是通过ID删除还是Name删除。
		1 ） 通过ID检索所有数据，如存在则返回所有数据，未找到返回查询不到对应数据。
		2 ） 通过Name检索所有数据，如存在则返回所有数据，未找到返回查询不到对应数据。
		3 ） 通过Contact检索所有数据，如存在则返回所有数据，未找到返回查询不到对应数据。
		4 ） 通过Address检索所有数据，如存在则返回所有数据，未找到返回查询不到对应数据。
	*/
	var Num string
	fmt.Println(`请输入要搜索的类型：
	1 ）输入用户ID查询数据
	2 ）输入用户名称查询数据
	3 ）输入联系电话查询数据
	4 ）输入邮箱地址查询数据
	`)
	fmt.Scan(&Num)
	switch Num {
	case "1":
		UserQueryPerfrom("ID")
	case "2":
		UserQueryPerfrom("Name")
	case "3":
		UserQueryPerfrom("Contact")
	case "4":
		UserQueryPerfrom("Address")
	default:
		fmt.Println("输入有误，回到上层。")
	}
}

func UserQueryPerfrom(elemet string) {
	/*
		根据传递进来的类型信息来确认通过什么类型来所有对应的数据
		UserData 用户输入查询的具体信息
		state 检索信息的状态返回码
			True 检测到具体的数据
			False 未检测到对应的数据
	*/
	var UserData, state string
	Users := make([]map[string]string, 0)
	fmt.Printf("请输入查询类型%s对应的信息：\n", elemet)
	fmt.Scan(&UserData)
	for _, v := range models.Users {
		// 判断如果遍历的数据与需要搜索的数据相同
		if v[elemet] == UserData {
			// 根据检索出的位置信息，将数据添加的切片中
			Users = append(Users, v)
			state = "True"
		}
	}
	// 如果遍历数据未找到对应数据，则返回删除失败
	if state == "True" {
		UsersList(&Users)
	} else {
		fmt.Printf("未找到%s为%s的数据，删除失败。\n", elemet, UserData)
	}

}
