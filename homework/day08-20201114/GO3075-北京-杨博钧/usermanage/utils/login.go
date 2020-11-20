package utils

import (
	"crypto/md5"
	"fmt"
	"os"
)

func Login() {
	chanse := 3
	fmt.Print("请输入用户名:")
	name := ""
	fmt.Scan(&name)
	i, exists := isUserExists(name)
	if exists {
		for {
			if chanse == 0 {
				fmt.Println("机会用尽，再见！")
				os.Exit(0)
			}
			fmt.Printf("当前还有%v次机会，请输入密码:", chanse)
			passwd := ""
			fmt.Scan(&passwd)
			if fmt.Sprintf("%x", md5.Sum([]byte(passwd))) == UsersList[i].Password {
				fmt.Println("登陆成功！")
				return
			} else {
				chanse--
			}
		}
	} else {
		fmt.Println("输入的用户不存在!退出...")
		os.Exit(0)
	}
}

