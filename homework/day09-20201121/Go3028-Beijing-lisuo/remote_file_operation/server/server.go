package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

// json := `{"cmd":"", "path":"", "fileName":""}`
// len(json)
//

var (
	addr     = "127.0.0.1:8889"
	servPath = "/tmp/"
)

type Response struct {
	Cmd      string `json:"cmd"`
	FilePath string `json:"filePath"`
	FileName string `json:"fileName"`
	Status   int    `json:"status"`
}

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server starts up, listenling: %s\n", addr)

	var d = make([]byte, 200)

	for {
		buf := bytes.NewBuffer(d)
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		toSend := Input("Server send: ")
		_, errW := conn.Write([]byte(toSend))
		if errW != nil {
			conn.Close()
			panic(errW)
		}
		_, errR := conn.Read(buf.Bytes())
		fmt.Println("From client: ", string(buf.Bytes()))
		buf.Reset()
		if errR != nil {
			conn.Close()
			panic(errR)
		}
	}

	fmt.Println("Server closed.")

}

func Input(s string) string {
	fmt.Print(s)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
