package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"usermanager_web/users"
)

var Udb users.Mydb
var Templatedir string

func Index(w http.ResponseWriter, r *http.Request) {
	tpl,err := template.ParseFiles(Templatedir+"/index.html")
	if err != nil {
		fmt.Println(err)
	}
	if err = tpl.ExecuteTemplate(w,"index.html",Udb.Getall()); err != nil {
		fmt.Println(err)
	}

}
//func Run(conn net.Conn, udb users.Mydb) {
//	var quit bool
//	p := 0
//	//验证3次密码
//	for p < 3 {
//		u, err := userutils.Readconn(conn)
//		if err != nil {
//			fmt.Printf("%s,断开连接...", err)
//			return
//		}
//		username := u.Userinfo["username"]
//		passwd := u.Userinfo["passwd"]
//		if udb.Auth(username, passwd) {
//			var sd = userutils.Msg{
//				Status: 200,
//				Ack:    true,
//				Data:   "欢迎进入用户管理系统",
//			}
//			err = userutils.Sendmes(conn, sd)
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//			var wsd = userutils.Msg{
//				Status: 200,
//				Ack:    true,
//				Data:   "1.添加用户.\n2.删除用户.\n3.修改用户.\n4.查找用户.\n5.退出\n",
//			}
//			err := userutils.Sendmes(conn, wsd)
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//			quit = true
//			break
//		} else {
//			var sd = userutils.Msg{
//				Status: 500,
//				Ack:    false,
//				Data:   "密码错误...",
//			}
//			err = userutils.Sendmes(conn, sd)
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//		}
//		p++
//	}
//	if p == 3 && quit == false {
//		fmt.Println("3次输入错误...")
//		return
//	}
//	for quit {
//
//		se, err := userutils.Readconn(conn)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		switch se.Code {
//		// 增加用户
//		case 1:
//			Add(se, udb, conn)
//		// 删除用户
//		case 2:
//			delUser(udb)
//		//// 修改用户
//		case 3:
//			modifyUser(udb)
//		// 查找用户
//		case 4:
//			queryuser(se, udb, conn)
//		// 退出
//		case 5:
//			fmt.Println("Bey...")
//			//同步到文件
//			err := udb.RotateSave()
//			if err != nil {
//				fmt.Println(err)
//			}
//			quit = false
//		default:
//			var sd = userutils.Msg{
//				Status: 500,
//				Ack:    false,
//				Data:   "未知选项",
//			}
//			err := userutils.Sendmes(conn, sd)
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//		}
//	}
//}
