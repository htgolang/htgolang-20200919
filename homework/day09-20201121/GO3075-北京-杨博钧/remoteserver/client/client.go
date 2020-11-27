package main

import (
	"fmt"
	"net"
	"remoteserver/model"
	"remoteserver/utils"
)

func main() {
	protocol := "tcp"
	addr := "127.0.0.1:8888"

	// 创建连接
	conn, err := net.Dial(protocol, addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	params := model.NewParams()

	request := utils.SerializeParams(params)

	conn.Write(request)

	utils.ClientWorker(conn, params)
}
