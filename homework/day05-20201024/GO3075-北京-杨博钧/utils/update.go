package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Update() {
	var name string
	fmt.Print("请输入要修改信息的人员名称:")
	fmt.Scan(&name)
	i, exists := isUserExists(name)
	if exists {
		LABEL:
		for {
			col := ""
			fmt.Print("请输入要修改的字段:")
			fmt.Scan(&col)
			col = strings.ToUpper(string(col[0])) + strings.ToLower(col[1:])
			var value string
			if strings.ToLower(col) != "password" {
				fmt.Print("请输入要修改的值:")
				fmt.Scan(&value)
			}
			switch col {
			case "Id" :
				data, err := strconv.Atoi(value)
				if err != nil {
					fmt.Println("修改值有误，请检查")
					continue LABEL
				}
				UsersList[i].Id = data
				break LABEL
			case "Addr" :
				UsersList[i].Addr = value
				break LABEL
			case "Name" :
				UsersList[i].Name = value
				break LABEL
			case "Tel" :
				UsersList[i].Tel = value
				break LABEL
			case "Birthday" :
				date, err := time.Parse("2006-01-02", value)
				if err != nil {
					fmt.Println("修改值有误，请检查")
					continue LABEL
				}
				UsersList[i].Birthday = date
				break LABEL
			case "Password" :
				pass := SetPasswd(UsersList[i].Name)
				UsersList[i].Password = pass
				break LABEL
			default :
				fmt.Println("输入字段异常，请检查")
				continue LABEL
			}
		}
		fmt.Printf("修改后用户表信息:\n%v", UsersList)
	} else {
		fmt.Println("人员不存在，请检查")
	}
}