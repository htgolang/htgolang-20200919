package utils

import (
	"fmt"
	"sort"
	"usermanage/model"
)

//删除函数
func DelUser(name string) {
	sort.Sort(UsersList)
	if i, exists := isUserExists(name); exists {
		if i == 0 && len(UsersList) == 1 {
			UsersList = []model.User{}
		} else if i == 0 && len(UsersList) > 1 {
			UsersList = UsersList[i + 1:]
		} else if i == len(UsersList) - 1 {
			UsersList = UsersList[:i]
		} else {
			UsersList = append(UsersList[:i], UsersList[i + 1:]...)
		}
	}
}
//根据用户名进行删除
func Del() {
	var name string
	fmt.Print("请输入要删除用户信息的姓名:")
	fmt.Scan(&name)
	i, exists := isUserExists(name)
	if exists {
		fmt.Printf("要删除以下用户的记录:\n%v\n", UsersList[i])
		choise := ""
		LABEL:
		for {
			fmt.Print("是否要进行删除操作(y/n)?:")
			fmt.Scan(&choise)
			switch choise {
			case "y", "Y":
				DelUser(name)
				break LABEL
			case "n", "N":
				break LABEL
			default:
				fmt.Println("输入错误，请重新输入")
			}
		}
		fmt.Printf("操作后用户列表数据为:\n%v", UsersList)
	} else {
		fmt.Println("要删除的用户不存在，请检查")
	}
}