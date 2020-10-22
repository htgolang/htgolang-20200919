package funcs

import (
	"crypto/md5"
	"fmt"
)

// 定义初始密码
var UserPasswd = strToMd5("123456")

func md5check() (Bool bool) {
	/*
		对比用户输入密码是否正确
			密码正确返回 true
			密码错误返回 false
	*/

	var Passwd string
	for i := 0; i < 3; i++ {
		if i == 0 {
			fmt.Println(`请输入密码：`)
			fmt.Scan(&Passwd)
		} else {
			fmt.Println(`密码输入错误，请重新输入密码：`)
			fmt.Scan(&Passwd)
		}

		// 对比输入值的md5是否与默认值的md5一致
		if strToMd5(Passwd) == UserPasswd {
			Bool = true
			return Bool
		}
	}
	fmt.Println("密码错误次数超过三次，已退出。Bay~，")
	return Bool
}

func strToMd5(txt string) (md5Value string) {
	md5Value = fmt.Sprintf("%X", md5.Sum([]byte(txt)))
	return
}
