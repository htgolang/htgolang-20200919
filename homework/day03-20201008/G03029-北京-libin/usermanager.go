package main

import (
	"fmt"
	"strings"
)

var users = []map[string]string{}
var mapFormat = [3]string{"username", "contact", "address"}

func main() {
	var name, contact, address string
	// fmt.Println
	for i := 0; i < 5; i++ {

		fmt.Scanln(&name, &contact, &address)
		fmt.Println(name, contact, address)
		Add(name, contact, address)

	}

	fmt.Println(users)
	Query()
	deleteSlice(2)
	fmt.Println(users)
}

func Add(name string, contact string, address string) {
	/*
			增加 add函数 从命令行分别输入名称、联系方式、通信地址
		生成ID => 查找users中最大的id+1（无元素id=1） 放入到users
	*/
	userMap := make(map[string]string)
	userMap["username"] = name
	userMap["contact"] = contact
	userMap["address"] = address

	users = append(users, userMap)

}

//查询
func Query() {
	var queryString string
	fmt.Println("请按顺序输入用户的名称，地址以及联系方式，用空格隔开")
	fmt.Scanln(&queryString)

	for userIndex, userInfo := range users {

		for _, v := range mapFormat {
			if strings.Contains(userInfo[v], queryString) {
				fmt.Println(users[userIndex])
			}
		}
	}
}

func deleteSlice(id int) []map[string]string {
	//根据id进行删除
	fmt.Println("请输入要删除的id对应的信息是,", users[id])
	var choice string
	fmt.Scanln(&choice)
	if choice == "yes" || choice == "y" || choice == "Y" || choice == "Yes" {
		users = append(users[:id], users[id+1:]...)
	}

	return users
}
