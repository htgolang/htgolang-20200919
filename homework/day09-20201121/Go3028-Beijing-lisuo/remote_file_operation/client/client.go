package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	addr        = "127.0.0.1:8889"
	defaultPath = "/tmp/"
)

type CommandBody struct {
	Cmd      string `json:"cmd"`
	FilePath string `json:"filePath"`
	FileName string `json:"fileName"`
	Status   int    `json:"status"`
}

func main() {
	var d = make([]byte, 20)
	for {
		buf := bytes.NewBuffer(d)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}
		_, errR := conn.Read(buf.Bytes())
		if errR != nil {
			conn.Close()
			panic(errR)
		}
		fmt.Printf("From server: %s\n", buf.Bytes())
		buf.Reset()

		toSend := Input("Client send: ")
		_, errW := conn.Write([]byte(toSend))
		if errW != nil {
			conn.Close()
			panic(errW)
		}
	}
	fmt.Println("Disconnected.")
}

func Input(s string) string {
	fmt.Print(s)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
