package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var (
	addr        = "suosuoli.cn:8889"
	defaultPath = "/tmp/"

	contentLenStr = 5
)

type CommandBody struct {
	Cmd      string `json:"cmd"`
	FilePath string `json:"filePath"`
	FileName string `json:"fileName"`
}

type ResponseBody struct {
	CommandBody
	Status int `json:"status"`
}

func main() {
	//for {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	cmd := CommandBody{"ls", "/tmp/", "server.log"}
	WriteContentLen(conn, cmd)
	WriteCmdBody(conn, cmd)
	resB := ReadResponseBody(conn)
	fmt.Println("resB: ", resB)

	// read cmd's result
	var buf = make([]byte, 200)
	_, errr := conn.Read(buf)
	if errr != nil {
		panic(errr)
	}
	fmt.Println("list: ", string(buf))
	conn.Close()

	//	}
	fmt.Println("Disconnected.")
}

func Input(s string) string {
	fmt.Print(s)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func WriteCmdBody(c net.Conn, cmdBody CommandBody) {
	b, _ := json.Marshal(cmdBody)
	_, errW := c.Write(b)
	if errW != nil {
		c.Close()
		panic(errW)
	}

}

func ReadResponseBody(c net.Conn) ResponseBody {
	conLen := ReadContentLen(c)
	var d = make([]byte, conLen)
	buf := bytes.NewBuffer(d)
	_, errR := c.Read(buf.Bytes())
	if errR != nil {
		c.Close()
		panic(errR)
	}
	responseBytes := buf.Bytes()
	var response = ResponseBody{}
	errUnmarshal := json.Unmarshal(responseBytes, &response)
	if errUnmarshal != nil {
		panic(errUnmarshal)
	}
	return response
}

func WriteContentLen(c net.Conn, cmd CommandBody) {
	bt, err := json.Marshal(cmd)
	if err != nil {
		c.Close()
		panic(err)
	}
	contentLen := len(string(bt))
	lenStr := fmt.Sprintf("%05d", contentLen)
	_, errW := c.Write([]byte(lenStr))
	if errW != nil {
		c.Close()
		panic(errW)
	}
}

func ReadContentLen(c net.Conn) int {
	var buf = make([]byte, contentLenStr)
	_, errRead := c.Read(buf)
	if errRead != nil {
		c.Close()
		panic(errRead)
	}
	len, err := strconv.Atoi(string(buf))
	if err != nil {
		panic(err)
	}
	return len
}
