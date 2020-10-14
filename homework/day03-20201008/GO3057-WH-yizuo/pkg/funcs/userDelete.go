package funcs

import (
	"GO3057-WH-yizuo/pkg/models"
	"fmt"
)

func UserDelete() {
	/*
	   根据ID或者用户名删除对应的用户
	   获取用户输入的值来判断是通过ID删除还是Name删除。
	   1 ） 通过ID检索所有数据，如存在删除并返回已删除，未找到返回删除失败。
	   2 ） 通过Name检索所有数据，如存在删除并返回已删除，未找到返回删除失败。
	*/
	var Num string
	fmt.Println(`请输入要删除的类型：
	1 ）输入用户ID删除
	2 ）输入用户名称删除
	`)
	fmt.Scan(&Num)
	switch Num {
	case "1":
		UserDeletePerform("ID")
	case "2":
		UserDeletePerform("Name")
	default:
		fmt.Println("输入有误，回到上层。")
	}
}

func UserDeletePerform(elemet string) {
	/*
		根据传递进来的类型信息来确认通过什么类型来删除对应的值
		UserData 用户输入删除的具体信息
		state 检索信息的状态返回码
			True 检测到具体的数据
			False 未检测到对应的数据
	*/
	var UserData, state string
	fmt.Printf("请输入要删除的%s信息：\n", elemet)
	fmt.Scan(&UserData)
	for k, v := range models.Users {
		// 判断如果遍历的数据与需要删除的数据相同
		if v[elemet] == UserData {
			// 根据检索出的位置信息，删除数据
			models.Users = append((models.Users)[:k], (models.Users)[k+1:]...)
			fmt.Printf("%s 为 %s 的数据已删除. \n", elemet, UserData)
			state = "True"
			return
		}
	}
	// 如果遍历数据未找到对应数据，则返回删除失败
	state = "False"
	if state == "False" {
		fmt.Printf("未找到%s为%s的数据，删除失败。\n", elemet, UserData)
	}

}
