package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func tablewriterfunc(userinfo []map[string]string) {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetAutoFormatHeaders(false)
	t.SetAutoWrapText(false)
	t.SetReflowDuringAutoWrap(false)

	t.SetHeader([]string{"ID", "名字", "电话号码", "联系地址"})
	for _, v := range userinfo {
		t.Append([]string{v["id"], v["name"], v["phone"], v["addr"]})
	}

	t.Render()
}

type UserManager struct {
	users []map[string]string
}

func (this *UserManager) addUser() {
	var (
		id    string
		name  string
		phone string
		addr  string
	)
	userInfo := map[string]string{
		"id":    "",
		"name":  "",
		"phone": "",
		"addr":  "",
	}

	fmt.Print("姓名：")
	fmt.Scan(&name)
	fmt.Print("\n联系方式(电话)：")
	fmt.Scan(&phone)
	fmt.Print("\n通信地址：")
	fmt.Scan(&addr)

	//判断用户格个数
	if len(this.users) == 0 {
		//如果切片中没有用户，则设置第一位用户id为1
		id = "1"
	} else {
		//第二为用户id为2，依次类推
		id = strconv.Itoa(len(this.users) + 1)
	}

	userInfo["id"] = id
	userInfo["name"] = name
	userInfo["phone"] = phone
	userInfo["addr"] = addr

	this.users = append(this.users, userInfo)

	// fmt.Println(this.users)
	tablewriterfunc(this.users)
}

func (this *UserManager) delUser() {
	var (
		id       string
		isDelete string
	)

	if len(this.users) == 0 {
		fmt.Println("空用户!!!")
		return
	}

	fmt.Print("请输入需要删除的用户id：")
	fmt.Scan(&id)

	for i, v := range this.users {
		if v["id"] == id {
			fmt.Println(v)
			fmt.Print("请输入y删除：")
			fmt.Scan(&isDelete)

			if isDelete == "y" && i == 0 {
				this.users = this.users[1:]
			} else if isDelete == "y" && i == len(this.users)-1 {
				this.users = this.users[:len(this.users)-1]
			} else if isDelete == "y" {
				this.users = append(this.users[:i], this.users[i+1:]...)
			}
		}
	}

	// fmt.Println(this.users)
	tablewriterfunc(this.users)
}

func (this *UserManager) modifyUser() {
	var (
		id       string
		name     string
		phone    string
		addr     string
		isModify string
	)

	fmt.Print("id：")
	fmt.Scan(&id)

	for _, v := range this.users {
		if v["id"] == id {
			fmt.Println(v)
			fmt.Print("请输入y修改：")
			fmt.Scan(&isModify)

			if isModify == "y" {
				fmt.Print("姓名：")
				fmt.Scan(&name)
				fmt.Print("\n联系方式(电话)：")
				fmt.Scan(&phone)
				fmt.Print("\n通信地址：")
				fmt.Scan(&addr)

				v["name"] = name
				v["phone"] = phone
				v["addr"] = addr

				fmt.Println(v)
			} else {
				fmt.Println("取消修改")
			}

		}
	}

	// fmt.Println(this.users)
	tablewriterfunc(this.users)
}

func (this *UserManager) queryUser() {
	var (
		info string
		flag bool
	)
	if len(this.users) == 0 {
		fmt.Println("空用户!!!")
		return
	}

	fmt.Print("请输入要查找的信息：")
	fmt.Scan(&info)
	for i, v := range this.users {
		flag = false
		for _, in := range v {
			if strings.Contains(in, info) && flag == false {
				flag = true
				fmt.Println(this.users[i])
			}
		}
	}
}

func (this *UserManager) MainMenu() {
	var isExist = false
	for {
		fmt.Println(`
---用户管理---
1、添加用户
2、删除用户
3、修改用户
4、查找用户
5、退出(q键)`)

		fmt.Printf("%v", this.users)
		var userchoice string
		fmt.Print("请选择：")
		fmt.Scan(&userchoice)
		switch {
		case userchoice == "1":
			this.addUser()
		case userchoice == "2":
			this.delUser()
		case userchoice == "3":
			this.modifyUser()
		case userchoice == "4":
			this.queryUser()
		case strings.ToLower(userchoice) == "q":
			isExist = true
		default:
			fmt.Println("请选择(1/2/3/4)")
		}

		if isExist {
			break
		}
	}
}

func MyUserManager() *UserManager {
	return &UserManager{
		users: []map[string]string{},
	}
}
