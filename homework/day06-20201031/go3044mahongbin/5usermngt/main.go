package main

import (
	"fmt"
	"go3044/modules"
	"go3044/utils"
)

var ops = map[string]func() int{
	"add":    modules.AddUser,
	"delete": modules.DelUser,
	"modify": modules.ModifyUser,
	"find":   modules.FindUser,
	"list":   modules.ListUser,
	"logout": modules.LogoutUser,
	"exit":   modules.ExitUser,
	"help":   modules.HelpUser,
}

func userAuth(tmpN, tmpP string) bool {
	authOK := false
	enP := utils.PasswordEncrypt(tmpP)
	for _, u := range modules.SliceU {
		if u.Name == tmpN && u.Password == enP {
			authOK = true
			break
		}
	}
	return authOK
}

func opsCheck() (x int) {
	x = 0
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("指令输入有误!请重试!")
			x = -1
		}
	}()

	x = ops[utils.Input(">>>>>请选择操作指令:")]()
	return x //当输入有误时,此代码不会被执行
}

func main() {
	// 测试代码
	// fmt.Println(utils.GenID(sliceUID))
	// fmt.Println(utils.Input("输入"))
	// fmt.Println(utils.PasswordEncrypt("123456"))
	//

	//正式流程
	// 1 初始化admin
	modules.InitAdmin()
	// 2 登陆验证
START:
	chances := 3
	for {
		if chances == 0 {
			fmt.Println("登录机会已耗尽!债见!")
			break
		}

		tmpN := utils.Input("请输入登录用户名:")
		tmpP := utils.Input("请输入登录 密码 :")
		authOK := userAuth(tmpN, tmpP)
		chances--

		if !authOK {
			fmt.Println("登录失败,请再次尝试!剩余机会:", chances)
			fmt.Println("--------------------------")
		} else {
			utils.CallClear()
			fmt.Println("*********登录成功!*********")
		AGAIN:
			for {
				switch opsCheck() {
				case -1:
					goto AGAIN
				case 0:
					goto START
				case 1:
					goto END
				}
			}
		}
	}
END:
}
