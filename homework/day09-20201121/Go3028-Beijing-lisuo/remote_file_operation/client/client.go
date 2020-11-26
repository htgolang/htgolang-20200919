package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	addr         = "suosuoli.cn:8889"
	defaultPath  = "/tmp/"
	headLen      = 5
	downloadPath = "/tmp/down/"
)

type CommandBody struct {
	Cmd      string `json:"cmd"`
	FilePath string `json:"filePath"`
	FileName string `json:"fileName"`
	FileSize int    `json:"fileSize"`
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

	cmd := CommandBody{"get", "/", "tmpfile.tail", 0}
	WriteHeadLen(conn, cmd)
	WriteHeadBody(conn, cmd)
	if cmd.Cmd == "put" {
		// conn.Write(filepath.Join(cmd.FilePath, cmd.FileName))
	}
	resB := ReadHeadBody(conn)
	fmt.Println("resB: ", resB)
	//HandleLS(conn, &resB)
	HandleGET(conn, &cmd, &resB)
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

// =========== protocol ===========
func WriteHeadLen(c net.Conn, cmd CommandBody) {
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

func WriteHeadBody(c net.Conn, cmdBody CommandBody) {
	b, _ := json.Marshal(cmdBody)
	_, errW := c.Write(b)
	if errW != nil {
		c.Close()
		panic(errW)
	}

}

func ReadHeadLen(c net.Conn) int {
	var buf = make([]byte, headLen)
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

func ReadHeadBody(c net.Conn) ResponseBody {
	conLen := ReadHeadLen(c)
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

// =========== data transfer ===========
func HandleLS(c net.Conn, cmd *ResponseBody) {
	res := ReadHeadBody(c)
	fileListLen := res.FileSize
	fmt.Println("file list len: ", fileListLen)

	var buf = make([]byte, fileListLen)
	_, err := c.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}

func HandleGET(c net.Conn, cmd *CommandBody, res *ResponseBody) {
	WriteHeadLen(c, *cmd)
	WriteHeadBody(c, *cmd)
	responseB := ReadHeadBody(c)
	fileSize := responseB.FileSize
	var buf = make([]byte, fileSize)
	_, err := c.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("file name: %v\n", responseB.FileName)
	fmt.Printf("file content: %v\n", string(buf))
	filePath := filepath.Join(cmd.FilePath, cmd.FileName)
	if err := os.MkdirAll(filePath, os.ModeDir); err != nil {
		panic(err)
	}
	f, errC := os.Create(filepath.Join(downloadPath, cmd.FileName))
	if errC != nil {
		panic(errC)
	}

	c.Read(buf)
	errW := ioutil.WriteFile(f.Name(), buf, os.ModeDir)
	if errW != nil {
		panic(errW)
	}

}
