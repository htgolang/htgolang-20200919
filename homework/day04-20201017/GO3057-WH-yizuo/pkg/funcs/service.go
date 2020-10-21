package funcs

import (
	"GO3057-WH-yizuo/pkg/models"
	"fmt"
	"strconv"
)

func FindLargestElementID() (ID string) {
	/*
	   查看现有数据中ID最大的值，将这个值+1作为我们新用户的ID并返回
	*/
	var num int
	for _, v := range models.Users {
		// 将搜索到的ID数据转换为int
		v, err := strconv.Atoi(v["ID"])
		if err != nil {
			return
		}
		// 对比大小将ID的值设置为最大的那个
		if num < v {
			num = v
		}
	}
	// 将最后找到的值转换为string并返回
	ID = strconv.Itoa(num + 1)
	return
}

func InuptUsersElement() (Name, Contact, Address string) {
	// 获取用户输入的 Name
	fmt.Println("请输入用户名称：")
	fmt.Scan(&Name)
	// 获取用户输入的 Contact
	fmt.Println("请输入手机号码：")
	fmt.Scan(&Contact)
	// 获取用户输入的 Address
	fmt.Println("请输入邮箱地址：")
	fmt.Scan(&Address)
	return
}
