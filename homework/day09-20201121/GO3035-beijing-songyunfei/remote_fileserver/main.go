package main

import (
	"fmt"
	"net"
	"remote_fileserver/fileserver"
)

func main() {
	fileserver.Basedir = "./basedir"
	addr := "127.0.0.1:8888"
	protocol := "tcp"
	listen, err := net.Listen(protocol, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("启动完成..")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%s,已连接\n", conn.RemoteAddr())
		go fileserver.Run(conn)

	}
}
