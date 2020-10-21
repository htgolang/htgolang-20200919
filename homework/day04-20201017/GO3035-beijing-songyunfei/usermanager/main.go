package main

import (
	"crypto/md5"
	"fmt"
	"usermanager/functions"
)

const adminpass  = "admin"
var pass string
var funcsmap map[string]func()
func init()  {
	pass = summd5(adminpass)
	funcsmap = make(map[string]func())
	funcsmap["add"] = functions.Useradd
	funcsmap["del"] = functions.DelUser
	funcsmap["modify"] = functions.ModifyUser
	funcsmap["query"] = functions.QueryUser
}

func summd5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return fmt.Sprintf("%X",ctx.Sum(nil))
}

func main()  {
	var quit bool
	p := 0
	for p <3{
		var inputpass string
		fmt.Printf("请输入密码:")
		_,_ = fmt.Scanln(&inputpass)
		if  mdpass := summd5(inputpass); mdpass == pass{
			quit = true
			fmt.Println("欢迎进入系统.")
			break
		}else {
			fmt.Println("密码错误....")
		}
		p++
	}
	if p == 3 && quit == false {
		fmt.Println("3次输入错误...")
	}
	for quit {

		var s int
		fmt.Printf("1.添加用户.\n2.删除用户.\n3.修改用户.\n4.查找用户.\n5.退出\n请输入序号(1~5):")
		_, _ = fmt.Scanln(&s)
		switch s {
		// 增加用户
		case 1:
			run,ok := funcsmap["add"]
			if ok  {
				run()
			}else {
				fmt.Println("Get func error")
			}
		// 删除用户
		case 2:
			run,ok := funcsmap["del"]
			if ok  {
				run()
			}else {
				fmt.Println("Get func error")
			}
		// 修改用户
		case 3:
			run,ok := funcsmap["modify"]
			if ok  {
				run()
			}else {
				fmt.Println("Get func error")
			}
		// 查找用户
		case 4:
			run,ok := funcsmap["query"]
			if ok  {
				run()
			}else {
				fmt.Println("Get func error")
			}
		// 退出
		case 5:
			fmt.Println("Bey...")
			quit = false

		default:
			fmt.Println("输入错误...")
		}
	}
}
