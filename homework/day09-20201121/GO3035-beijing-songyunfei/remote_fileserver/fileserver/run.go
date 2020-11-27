package fileserver

import (
	"fmt"
	"net"
	pubfunc "remote_fileserver/pubrw"
)

//Basedir 根路径
var Basedir string

//Run 启动入口
func Run(conn net.Conn) {
	for {
		cmdsl, err := pubfunc.Readdata(conn)
		if err != nil {
			fmt.Printf("%s,断开连接.\n",conn.RemoteAddr())
			return
		}
		cmd, err := pubfunc.Unmarshalcmd(cmdsl)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch cmd.Cmd {
		case "ls":
			ls(cmd, conn)
		case "rm":
			deletefile(cmd, conn)
		case "put":
			err := putfile(cmd, conn)
			if err != nil {
				fmt.Println(err)
				return
			}
		case "get":
			err = sendfile(cmd, conn)
			if err != nil {
				fmt.Println(err)
				return
			}
		case "quit":
			break

		default:
			fmt.Println("unknown arg")
			break

		}
	}

}
