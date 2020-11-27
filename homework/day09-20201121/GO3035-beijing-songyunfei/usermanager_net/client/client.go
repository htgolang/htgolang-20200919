package main

import (
	"fmt"
	"net"
	"usermanager_net/userutils"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	var cmd = userutils.Cmd{
		Cmd:      "auth",
		Code:     100,
		Userinfo: make(map[string]string),
	}
	for {
		cmd.Userinfo["username"] = userutils.Input("用户名:")
		cmd.Userinfo["passwd"] = userutils.Input("密码:")
		err = userutils.SendCmd(conn, cmd)
		if err != nil {
			fmt.Println(err)
			return
		}
		msg, err := userutils.ReadMsg(conn)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(msg.Data)
		if msg.Ack {
			s, err := userutils.ReadMsg(conn)
			if err != nil {
				fmt.Println(err)
				return
			}
			for {
				fmt.Printf("%s", s.Data)
				c := userutils.Input("请输入序号(1~5):")
				switch c {
				case "1":
					var cmd = userutils.Cmd{
						Cmd:      "add",
						Code:     1,
						Userinfo: make(map[string]string),
					}
					cmd.Userinfo["username"] = userutils.Input("请输入用户名:")
					cmd.Userinfo["addr"] = userutils.Input("请输入地址:")
					cmd.Userinfo["tel"] = userutils.Input("请输入联系电话:")
					cmd.Userinfo["passwd"] = userutils.Input("请输入密码:")
					cmd.Userinfo["brth"] = userutils.Input("请输入生日(示例:1999-01-05):")
					header := []string{"用户名", "地址", "联系电话", "生日"}
					data := [][]string{{cmd.Userinfo["username"], cmd.Userinfo["addr"], cmd.Userinfo["tel"], cmd.Userinfo["brth"]}}
					userutils.Showintable(header, data)
					yorn := userutils.Input("确认添加y/n:")
					if yorn == "y" {
						if err := userutils.SendCmd(conn, cmd); err != nil {
							fmt.Println(err)
						}
						ames, err := userutils.ReadMsg(conn)
						if err != nil {
							fmt.Println(err)
						}
						fmt.Println(ames.Data)
					}
				case "4":
					qs := userutils.Input("请出入关键字:")
					var c = userutils.Cmd{
						Cmd:     "4",
						Code:     4,
						Data:     qs,
						Userinfo: nil,
					}
					err = userutils.SendCmd(conn,c)
					if err != nil {
						fmt.Println(err)
						return
					}
					qmes,err := userutils.ReadMsg(conn)
					if err != nil {
						fmt.Println(err)
						return
					}
					if qmes.Ack{
						fmt.Println(qmes.Data)
					}

				}
			}

		}
	}

}