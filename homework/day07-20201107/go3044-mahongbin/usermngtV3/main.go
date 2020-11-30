package main

import (
	"flag"
	"fmt"
	"usermanage/modules"
	"usermanage/utils"
)

var ops = map[string]func() int{
	"add":    modules.AddUser,
	"delete": modules.DelUser,
	"modify": modules.ModifyUser,
	"find":   modules.FindUser,
	"list":   modules.ListUser,
	"exit":   modules.ExitUser,
	"help":   modules.HelpUser,
}

//opsCheck 登陆之后的操作选择
func opsCheck() (x int) {
	x = 0

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("输入指令无法识别!!!请输入'help'查看帮助信息")
			x = -1
		}
	}()

	x = ops[utils.Input(">>>>>请选择操作指令:")]()
	return x //当输入有误时,此代码不会被执行
}

func main() {

	//正式流程
	// 1 初始化admin

	//数据库文件格式初始化
	var initType, dbPath string
	var help bool
	flag.StringVar(&initType, "init", "", "init user DB file ... gob|csv|json")
	flag.StringVar(&dbPath, "load", "", "load user DB File")
	flag.BoolVar(&help, "h", false, "Help Info")

	modules.DbFileInitType = initType

	flag.Usage = func() {
		fmt.Println(`
Usage:
	usermngt -init [gob|csv|json]
OR	usermngt -load userDB.[gob|csv|json]
Options:`)
		flag.PrintDefaults()
	}

	flag.Parse()

	// fmt.Println(initType, dbPath, help) //debug

	if help {
		flag.Usage()
		return
	}

	switch {
	case dbPath != "" && initType != "":
		fmt.Println("不能同时使用init和load!!!")
		return
	case dbPath != "" && initType == "":
		err := modules.LoadDbFile(dbPath)
		if err != nil {
			fmt.Println(err)
			return
		}
	case dbPath == "" && initType != "":
		err := modules.SetDbNameAndMode(initType)
		if err != nil {
			fmt.Println(err)
			return
		}
		modules.InitAdmin()
		modules.SaveDbFile("." + initType)
	default: //dbPath == "" && initType == "":
		fmt.Println("输入参数为空...")
		return
	}

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
		authOK := modules.UserAuth(tmpN, tmpP)
		chances--

		if !authOK {
			fmt.Println("登录失败,请再次尝试!剩余机会:", chances)
			fmt.Println("--------------------------")
		} else {
			// utils.CallClear()
			fmt.Println("*********登录成功!*********")
		AGAIN:
			for {
				//opsCheck 登陆之后
				switch opsCheck() {
				case -1:
					goto AGAIN //add delete modify ...
				case 0:
					goto START //logout
				default:
					modules.SaveOrNot()
					goto END //exit
				}
			}
		}
	}
END:
}
