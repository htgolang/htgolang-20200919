package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"usermanagement/modules"
	"usermanagement/utils"
)

func Login(m *modules.UserManager) bool {
	fmt.Println("Please login!!")
	username := utils.Input("Please input username: ")
	passwd := utils.Md5Text(utils.Input("Please input password: "))

	for _, user := range m.Users {
		if user.Name == username && user.Password == passwd {
			fmt.Println("login success!!")
			return true
		}
	}
	fmt.Println("username or password is incorrect..")
	return false

}

func DeleteUserById(m *modules.UserManager) {
	u_id := utils.Input("Please input user id: ")
	for idx, user := range m.Users {
		if uid, err := strconv.Atoi(u_id); err == nil {
			if user.Id == uid {
				fmt.Println("User information to be delete.")
				fmt.Printf("%#v\n", user)
				confirm := utils.Input("Confirm delete?(Y/m)")
				if strings.ToLower(confirm) == "y" || strings.ToLower(confirm) == "yes" {
					m.DeleteUser(idx)
					fmt.Println("[-]delete success")
					break
				}
			}
		}
	}
}

func AddUser(m *modules.UserManager) {
	name := utils.Input("Please input username: ")
	address := utils.Input("Please input address: ")
	tel := utils.Input("Please input tel: ")
	birthday := utils.TimeConversion(utils.Input("Please input your birthday(Y-m-d): "))
	passwd := utils.Md5Text(utils.Input("Please input password: "))
	if m.FindUserByName(name) {
		fmt.Printf("[+] %s already exists\n", name)
		return

	}
	m.AddUser(&modules.UserInfo{
		Id:       m.GetId(),
		Name:     name,
		Address:  address,
		Tel:      tel,
		Birthday: birthday,
		Password: passwd,
	})
	fmt.Printf("[+]%s add success!\n", name)
}

func QueryUser(m *modules.UserManager) {
	name := utils.Input("Please inout username: ")
	user := m.QueryUser(name)
	if user == nil {
		fmt.Println("User does not exist")
		return
	}
	fmt.Printf("%#v\n", user)
}

func ModifyUser(m *modules.UserManager) {
	id := utils.Input("Please input id: ")
	uid, _ := strconv.Atoi(id)
	idx, isExist := m.FindUserById(uid)
	if isExist {
		fmt.Println("User information to be modified.")
		fmt.Printf("%#v\n", m.Users[idx])
		confirm := utils.Input("Confirm modify?(Y/m)")
		if strings.ToLower(confirm) == "y" || strings.ToLower(confirm) == "yes" {
			m.ModifyUser(idx, &modules.UserInfo{
				Id:       uid,
				Name:     utils.Input("Please input name: "),
				Address:  utils.Input("Please input address: "),
				Tel:      utils.Input("Please input tel: "),
				Birthday: utils.TimeConversion(utils.Input("Please input your birthday(Y-m-d): ")),
				Password: utils.Md5Text(utils.Input("Please input password: ")),
			})
			fmt.Println("modify sunccess!")
			return
		}
	}
	fmt.Println("User does not exist")
}
