package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	addr := "127.0.0.1:8888"
	protocol := "tcp"

	start := time.Now()
	conn, _ := net.Dial(protocol, addr)

	reader := bufio.NewReader(conn)
	fmt.Println(reader.ReadString('\n'))
	conn.Close()

	fmt.Println(time.Now().Sub(start))
}
