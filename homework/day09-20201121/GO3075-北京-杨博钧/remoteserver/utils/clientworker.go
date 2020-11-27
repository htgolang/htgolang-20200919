package utils

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"remoteserver/model"
)

// 根据不同参数做不同的处理
func ClientWorker(conn net.Conn, params *model.Params) {
	defer conn.Close()
	var code int
	var message string
	if params.Cmd == "ls" || params.Cmd == "rm" {
		code, message = GetStatus(conn)
	} else if params.Cmd == "get" {
		code, message = DownloadFile(conn, params.Dest)
	} else if params.Cmd == "put" {
		code, message = UploadFile(conn, params.Src)
	}
	fmt.Printf("code: %d\tmessage: %s", code, message)
}

func GetStatus(conn net.Conn) (int, string) {
	reader := bufio.NewReader(conn)
	ctx, _ := reader.ReadBytes('\n')

	returnvalue := DeserializeReturns(ctx)

	return returnvalue.Code, returnvalue.Message
}

func DownloadFile(conn net.Conn, dest string) (int, string) {
	code, message := GetStatus(conn)
	if code != 200 {
		return code, message
	}

	reader := bufio.NewReader(conn)

	destfile, err := os.OpenFile(dest, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return 400, fmt.Sprint(err)
	}

	writer := bufio.NewWriter(destfile)
	defer writer.Flush()

	io.Copy(writer, reader)
	return 200, fmt.Sprintf("拷贝远程服务器到本地%v成功\n", dest)
}

func UploadFile(conn net.Conn, src string) (int, string) {
	srcfile, err := os.Open(src)
	if err != nil {
		return 400, fmt.Sprint(err)
	}
	defer srcfile.Close()
	reader := bufio.NewReader(srcfile)
	writer := bufio.NewWriter(conn)
	defer writer.Flush()
	io.Copy(writer, reader)

	return 200, "ok"
}