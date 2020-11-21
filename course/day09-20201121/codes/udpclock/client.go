package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addr := "127.0.0.1:8888"
	protocol := "udp"

	conn, err := net.Dial(protocol, addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	n, err := conn.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
	fmt.Println(n, err)
	ctx := make([]byte, 1024)
	n, err = conn.Read(ctx)
	fmt.Println(string(ctx[:n]))

	conn.Close()
}
