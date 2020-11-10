package funcs

import (
	"fmt"
	"homework/pkg/models"
	"strconv"
)
//查找现有最大的ID
func findMaxID() (ID string) {
	var num int
	for _,v:=range models.Users {
		v,err:=strconv.Atoi(v["ID"])
		if err!=nil {
			return
		}
		if num < v {
			num=v
		}
	}
	//将最大的ID+1后转为string
	ID=strconv.Itoa(num+1)
	return
}

func inputUser() (Name,Contact,Address string) {
	fmt.Println("请输入用户名：")
	fmt.Scan(&Name)
	fmt.Println("请输入联系方式：")
	fmt.Scan(&Contact)
	fmt.Println("请输入地址：")
	fmt.Scan(&Address)
	return
}