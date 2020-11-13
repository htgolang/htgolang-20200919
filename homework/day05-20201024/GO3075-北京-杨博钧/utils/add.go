package utils

import (
	"fmt"
	"time"
	"usermanage/model"
)

func Add() {
	fmt.Println("请输入要添加的用户信息")
	id := GetMaxId() + 1
	name := ""
	fmt.Print("请输入用户姓名:")
	fmt.Scan(&name)
	addr := ""
	fmt.Printf("请输入%v的联系地址:", name)
	fmt.Scan(&addr)
	tel := ""
	fmt.Printf("请输入%v的手机号码:", name)
	fmt.Scan(&tel)
	birthday := ""
	fmt.Printf("请输入%v的生日，请按照中2006-01-02格式输入日期:", name)
	fmt.Scan(&birthday)
	time, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		fmt.Println(err)
		fmt.Println("生日格式不正确无法解析")
		return
	}
	passwd := SetPasswd(name)
	user := model.User{id, name, addr, tel, time, passwd}
	if _, exists := isUserExists(name); exists {
		fmt.Println("用户已经存在，不再进行添加")
		return
	}
	UsersList = append(UsersList, user)
	fmt.Printf("操作后用户列表数据为:\n%v", UsersList)
}
