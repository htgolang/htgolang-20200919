package main

import (
	"fmt"
	"strconv"
	"strings"
)

var users = []map[string]string{} // 切片类型的map

func getId() string {
	id := 0
	for _, user := range users {
		if uid, err := strconv.Atoi(user["id"]); err == nil {
			if uid > id {
				id = uid
			}
		}
	}
	return strconv.Itoa(id + 1)
}

func input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return text
}

func add() {
	var (
		name string
		addr string
		tel  string
	)
	fmt.Printf("用户管理系统\n")
	for {
		fmt.Printf("输入名称：")
		fmt.Scan(&name)
		fmt.Printf("联系方式：")
		fmt.Scan(&addr)
		fmt.Printf("电话:")
		fmt.Scan(&tel)
		user := map[string]string{
			"id":   getId(),
			"name": name,
			"addr": addr,
			"tel":  tel,
		}

		users = append(users, user)
		fmt.Printf("当前用户为\n")
		fmt.Println(users)

		// 提示框输入函数
		var choice string
		fmt.Printf("退出添加用户y,继续添加输入n:")

		fmt.Scan(&choice)

		if choice == "y" {
			fmt.Printf("用户即将退出")
			fmt.Println(users)
			break
		}

	}

}

func findUserByid(id string) map[string]string {
	for _, user := range users {
		if user["id"] == id {
			return user
		}
	}
	return nil
}

func deleteUserByid(id string) {
	tempUsers := make([]map[string]string, 0, len(users)-1)
	for _, user := range users {
		if user["id"] != id {
			tempUsers = append(tempUsers, user)
		}
	}
	users = tempUsers
}

func deleteUser() {
	fmt.Print("请输入用户需要删除的id:")
	var delId string
	var choice string
	fmt.Scan(&delId)
	user := findUserByid(delId)
	if user == nil {
		fmt.Println("用户信息不存在")
	} else {
		fmt.Println("需要删除的用户信息如下：")
		fmt.Println(user)
		fmt.Printf("确定删除的吗?(Y/n)")
		fmt.Scan(&choice)
		if strings.ToLower(choice) == "y" || strings.ToLower(choice) == "yes" {
			deleteUserByid(delId)
		}

	}
}

func modifyUserByid(user map[string]string, id string) {
	for idx, tuser := range users {
		if tuser["id"] == "id" {
			users[idx] = user
			fmt.Println("idx",idx)
			break
		}
	}


}

func modify() {
	id :=input("请输入要修改的用户id:")
	user := findUserByid(id)
	if user == nil {
		fmt.Println("用户信息不存在")
	} else {
		fmt.Println("需要修改的用户信息如下：")
		fmt.Println(user)
		confirm :=input("确定要修改吗?(y/n)")
		if strings.ToLower(confirm) == "y" || strings.ToLower(confirm) == "yes" {
			user := map[string]string{
				"id":    id,
				"name":  input("name: "),
				"addr:": input("addr: "),
				"tel":   input("tel: "),
			}

			tempUsers :=make([]map[string]string,0,len(users))

			for _,tuser :=range users{
            if tuser["id"]==id{
            	tempUsers=append(tempUsers,user)
			}else {
				tempUsers=append(tempUsers,tuser)
			}
			}
			users=tempUsers

		}
	}

}

func filter (user map[string]string,q string) bool  {
	return strings.Contains(user["name"],q)||
		strings.Contains(user["addr"],q)||
		strings.Contains(user["tel"],q)
}

func query() {
	q :=input("请输入查询信息:")
	fmt.Println("查询结果")
	for _,user :=range users{
		if filter(user,q){
			fmt.Println(user)
		}
	}

}

func main() {

	operates := map[string]func(){
		"add":    add,
		"modify": modify,
		"delete": deleteUser,
		"query":  query,
	}
	for {
		var text string
		fmt.Print("请输入指令: ")
		fmt.Scan(&text)
		if text == "exit" {
			fmt.Println("退出")
			break
		}
		if op, ok := operates[text]; ok {
			op()
		} else {
			fmt.Println("指令错误")
		}

	}






}
