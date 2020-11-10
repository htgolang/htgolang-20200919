package funcs

import (
	"crypto/md5"
	"fmt"
)

//初始密码
var UserDefautPwd = toMd5("admin123")

func md5Check() (Bool bool) {
	//检查用户输入的密码是否正确,错误3次就退出
	var Pwd string
	for i:=0;i<3;i++ {
		if i==0 {
			fmt.Println("请输入您的密码：")
			fmt.Scan(&Pwd)
		}	else {
			fmt.Println("密码输入错误，请重新输入：")
			fmt.Scan(&Pwd)
		}
		if toMd5(Pwd)== UserDefautPwd {
			Bool = true
			return Bool
		}	
	}
	fmt.Println("密码输错3次。")
	return Bool
}

func toMd5(text string) (md5Value string) {
	md5Value = fmt.Sprintf("%X",md5.Sum([]byte(text)))
	return
}