package utils

import (
	"crypto/md5"
	"fmt"
)
//定义设置密码函数
func SetPasswd(name string) string {
	for {
		//对两次输入的密码进行校验，两次一致时再进行md5转换存储
		fmt.Printf("请输入%v账户新的密码:", name)
		passwd1 := ""
		fmt.Scan(&passwd1)
		fmt.Printf("请再次输入确认:")
		passwd2 := ""
		fmt.Scan(&passwd2)
		if passwd1 == passwd2 {
			return fmt.Sprintf("%x", md5.Sum([]byte(passwd1)))
			break
		} else {
			fmt.Println("检测到两次密码不一致，请重新设置密码")
			continue
		}
	}
	return ""
}