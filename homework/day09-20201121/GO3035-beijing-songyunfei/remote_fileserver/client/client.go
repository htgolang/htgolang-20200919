package main

import (
	"encoding/json"
	"fmt"
	"net"
	"remote_fileserver/data_mod"
	pubfunc "remote_fileserver/pubrw"
)

func main() {
	addr := "0.0.0.0:8888"
	pro := "tcp"
	conn, err := net.Dial(pro, addr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("连接服务器成功..")
	fmt.Printf("ls [arg],显示目录\nrm [arg],删除.\nput [file],上传文件.\nget [remoteFile],下载文件到本地.\nquit, 退出.\n")
	for {
		c := pubfunc.Input("请输入命令:")
		switch c.Cmd {
		case "ls":
			lsAndrm(conn, c)
		case "rm":
			lsAndrm(conn, c)
		case "put":
			err = pubfunc.Sendfile(conn, c)
			if err != nil {
				fmt.Println(err)
				return
			}
		case "get":
			err := getfile(c, conn)
			if err != nil {
				fmt.Println(err)
				return
			}
		case "quit":
			d,err := json.Marshal(c)
			if err != nil {
				fmt.Println(err)
				return
			}
			_ = pubfunc.Senddata(conn,d)
			return
		}

	}

}

func lsAndrm(conn net.Conn, c data_mod.Cmd) {
	sd, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	err = pubfunc.Senddata(conn, sd)
	rdata, err := pubfunc.Readdata(conn)
	data, err := pubfunc.Unmarshaldata(rdata)
	fmt.Println(string(data.Data))
}

func getfile(cmd data_mod.Cmd, conn net.Conn) error {
	sd, err := json.Marshal(cmd)
	if err != nil {
		fmt.Println(err)
	}
	err = pubfunc.Senddata(conn, sd)
	if err != nil {
		return err
	}
	err = pubfunc.Getfile(cmd, conn)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
