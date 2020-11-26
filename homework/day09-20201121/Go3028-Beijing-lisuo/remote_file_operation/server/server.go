package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

/*
cmds: ls, put, get, rm
path: /tmp/xxx
file:
*/

var (
	addr          = ":8889"
	servPath      = "/tmp/"
	contentLenStr = 5
)

type ResponseBody struct {
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

	//for {
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	rec := ReadCmdBody(conn)
	fmt.Println("rec: ", rec)
	rec.Status = 200
	WriteContentLen(conn, rec)
	WriteRespondBody(conn, rec)

	// parse cmdBody and decide to send file or list files
	ListFiles(conn, rec)

	conn.Close()
	fmt.Println("Server closed.")

}

func Input(s string) string {
	fmt.Print(s)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func WriteContentLen(c net.Conn, response ResponseBody) {
	bt, err := json.Marshal(response)
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

func ReadCmdBody(c net.Conn) ResponseBody {
	conLen := ReadContentLen(c)
	var d = make([]byte, conLen)
	buf := bytes.NewBuffer(d)
	_, errR := c.Read(buf.Bytes())
	if errR != nil {
		c.Close()
		panic(errR)
	}
	cmdBodyBytes := buf.Bytes()
	var resBeforeSend = ResponseBody{}
	errUnmarshal := json.Unmarshal(cmdBodyBytes, &resBeforeSend)
	if errUnmarshal != nil {
		panic(errUnmarshal)
	}
	return resBeforeSend
}

func WriteRespondBody(c net.Conn, resBody ResponseBody) {
	b, errMar := json.Marshal(resBody)
	if errMar != nil {
		panic(errMar)
	}
	_, errW := c.Write(b)
	if errW != nil {
		c.Close()
		panic(errW)
	}

}

func ListFiles(c net.Conn, res ResponseBody) {
	if res.Cmd == "ls" {
		c.Write([]byte(fmt.Sprintf("file %v is here.\n", filepath.Join(res.FilePath, res.FileName))))
	}
}
