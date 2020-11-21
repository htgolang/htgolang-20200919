package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addr := ":8888"
	protocol := "udp"

	packetConn, err := net.ListenPacket(protocol, addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 处理客户端
	for {
		ctx := make([]byte, 1024)
		n, addr, err := packetConn.ReadFrom(ctx)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("客户端[%s]发送数据: %s\n", addr, string(ctx[:n]))
		fmt.Println(packetConn.WriteTo([]byte(time.Now().Format("2006/01/02 15:04:05")), addr))
	}
	packetConn.Close()

}
