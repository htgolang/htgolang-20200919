package main

import (
	"bufio"
	"fmt"
	_ "io"
	"net"
	"remoteserver/utils"
)

func main() {

	addr := "0.0.0.0:8888"
	protocol := "tcp"

	// 启动监听服务
	listener, err := net.Listen(protocol, addr)
	if err != nil {
		fmt.Println(err)
		return

	}
	defer listener.Close()
	for {
		// 开启连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go ResloveRequest(conn)
	}

}

func ResloveRequest(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	request, _ := reader.ReadBytes('\n')
	params := utils.DeserializeParams(request)

	utils.ServerWorker(conn, params)
}