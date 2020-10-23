package function

import (
	"fmt"
	"strings"
)

const (
	modifyPrompts   string = "Are you sure modify (y/n): "
	idModifyPrompts string = "enter you want modify user is id: "
)

//Modify sss
func Modify() {
	usermap, userIndex := getUserMessage(idModifyPrompts)
	if len(userIndex) == 0 {
		fmt.Println("user not exist!!")
	} else if len(userIndex) == 1 {
		printTable(usermap)
		modifyUsers(userIndex[0], usermap)
	} else {
		fmt.Println("enter user is primary id : ")
	}
}

func modifyUsers(index int, user []map[string]string) {

	modifyUserInputMessage := []map[string]string{}
	//输入要修改的信息并打印
	m := getUserInputMessage()
	modifyUserInputMessage = append(modifyUserInputMessage, m)
	printTable(modifyUserInputMessage)

	//确认是否修改
	yn := getQueryUserInput(modifyPrompts)
	if strings.ToLower(yn) == "y" {
		fmt.Println(user)
		Users[index]["id"] = user[0]["id"]
		Users[index]["name"] = m["name"]
		Users[index]["tel"] = m["tel"]
		Users[index]["addr"] = m["addr"]
	}
}
