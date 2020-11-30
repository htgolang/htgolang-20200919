package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strconv"
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

// CommandBody save command
type CommandBody struct {
	Cmd      string `json:"cmd"`
	FilePath string `json:"filePath"`
	FileName string `json:"fileName"`
	FileSize int64  `json:"fileSize"`
}

// ResponseBody wrap a status attribute
type ResponseBody struct {
	CommandBody
	Status int `json:"status"`
}

func main() {
	//for {
	cmd := ParseCmd()
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	switch cmd.Cmd {
	case "ls":
		HandleLS(conn, &cmd)
	case "get":
		rStatus := HandleGET(conn, &cmd)
		if err := HandleError(rStatus); err != nil {
			fmt.Println(err)
		}
	case "put":
		resB := HandlePUT(conn, &cmd)
		fmt.Printf("resB: %#v\n", resB)
	case "rm":
		resB := HandleRM(conn, &cmd)
		fmt.Println("resB: ", resB)
	}

	conn.Close()

	//	}
	fmt.Println("Disconnected.")
}

// =========== protocol ===========

// WriteHead wrap WriteHeadLen and WriteHeadBody
func WriteHead(c net.Conn, cmd CommandBody) {
	WriteHeadLen(c, cmd)
	WriteHeadBody(c, cmd)
}

// WriteHeadLen send json head len to client
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

// WriteHeadBody send json head to server
func WriteHeadBody(c net.Conn, cmdBody CommandBody) {
	b, _ := json.Marshal(cmdBody)
	_, errW := c.Write(b)
	if errW != nil {
		c.Close()
		panic(errW)
	}

}

func readHeadLen(c net.Conn) int {
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

// ReadHeadBody read json response head from server
func ReadHeadBody(c net.Conn) ResponseBody {
	conLen := readHeadLen(c)
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

// HandleLS handles the ls command
func HandleLS(c net.Conn, cmd *CommandBody) {
	fmt.Println("cmd: ", *cmd)
	WriteHead(c, *cmd)
	response := ReadHeadBody(c)
	if err := HandleError(response.Status); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileListLen := response.FileSize
	fmt.Println("file list len: ", fileListLen)

	var buf = make([]byte, fileListLen)
	_, err := c.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}

// HandleGET handles the get command
func HandleGET(c net.Conn, cmd *CommandBody) int {
	WriteHead(c, *cmd)
	responseB := ReadHeadBody(c)
	if responseB.Status != 200 {
		return responseB.Status
	}
	fileSize := responseB.FileSize
	var buf = make([]byte, fileSize)
	_, err := c.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("file name: %v\n", responseB.FileName)
	//fmt.Printf("file content: %v\n", string(buf))
	//filePath := filepath.Join(cmd.FilePath, cmd.FileName)
	//if err := os.MkdirAll(cmd.FilePath, os.ModeDir); err != nil {
	//	panic(err)
	//}
	f, errC := os.Create(filepath.Join(downloadPath, cmd.FileName))
	if errC != nil {
		panic(errC)
	}

	//c.Read(buf)
	errW := ioutil.WriteFile(f.Name(), buf, os.ModeDir)
	if errW != nil {
		panic(errW)
	}
	fmt.Printf("ResponseBody: %#v\n", responseB)
	return responseB.Status
}

// HandlePUT handles put command
func HandlePUT(c net.Conn, cmd *CommandBody) ResponseBody {
	absFile := filepath.Join(cmd.FilePath, cmd.FileName)
	f, err := os.Open(absFile)
	if err != nil {
		panic(err)
	}
	fStat, errS := os.Stat(absFile)
	if errS != nil {
		panic(errS)
	}
	cmd.FileSize = fStat.Size()
	WriteHead(c, *cmd)
	buf := make([]byte, cmd.FileSize)
	reader := bufio.NewReader(f)
	reader.Read(buf)
	c.Write(buf)
	return ReadHeadBody(c)

}

// HandleRM handles rm command
func HandleRM(c net.Conn, cmd *CommandBody) ResponseBody {
	WriteHead(c, *cmd)
	return ReadHeadBody(c)
}

// =========== util ===========

// ParseCmd parse user cmd and fill HeadBody
func ParseCmd() CommandBody {
	/*
		type CommandBody struct {
			Cmd      string `json:"cmd"`
			FilePath string `json:"filePath"`
			FileName string `json:"fileName"`
			FileSize int    `json:"fileSize"`
		}
	*/
	cmd := flag.String("c", "", "specify the command(ls,get,put,rm)")
	filePath := flag.String("fp", "/", "specify the file path")
	fileName := flag.String("fn", "", "specify the file name")
	flag.Parse()
	if *cmd == "" {
		fmt.Println(errors.New("must specify a command(-c)"))
		flag.Usage()
		os.Exit(1)
	}
	cmdbody := CommandBody{
		Cmd:      *cmd,
		FilePath: *filePath,
		FileName: *fileName,
	}
	return cmdbody
}

//HandleError handles the errors returned from server
func HandleError(s int) error {
	switch s {
	case 200:
		return nil
	case 404:
		return errors.New("Server: can not find the file you asked")
	case 500:
		return errors.New("Server: server fault")
	default:
		return errors.New("Nodbody: something wrong")
	}
}
