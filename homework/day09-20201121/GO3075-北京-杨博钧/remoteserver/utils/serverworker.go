package utils

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"remoteserver/model"
)

func ServerWorker(conn net.Conn, params *model.Params) {
	if params.Cmd == "ls" {
		ListFile(params.Path, conn)
	} else if params.Cmd == "rm" {
		RemoveFile(params.Path, conn)
	} else if params.Cmd == "get" {
		PutFile(conn, params.Src)
	} else if params.Cmd == "put" {
		GetFile(conn, params.Dest)
	}
}

func ListFile(path string, conn net.Conn) {
	var returnvalue model.ReturnValue
	stat, err := os.Stat(path)
	if err != nil {
		returnvalue = model.ReturnValue{
			400,
			fmt.Sprint(err),
		}
	} else {
		returnvalue = model.ReturnValue{
		200,
		fmt.Sprintf("%v\t%v\t%v\t%v\n", stat.Mode(), stat.Size(), stat.ModTime(), stat.Name()),
	}
	}
	info := SerializeReturns(returnvalue)
	conn.Write(info)
}

func RemoveFile(path string, conn net.Conn) {
	var returnvalue model.ReturnValue
	err := os.RemoveAll(path)
	if err != nil {
		returnvalue = model.ReturnValue{
			400,
			fmt.Sprint(err),
		}
	}
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		returnvalue = model.ReturnValue{
			200,
			fmt.Sprintf("成功移除%v\n", path),
		}
	} else {
		returnvalue = model.ReturnValue{
			400,
			fmt.Sprint(err),
		}
	}
	info := SerializeReturns(returnvalue)
	conn.Write(info)
}

func PutFile(conn net.Conn,src string) {
	var returnvalue model.ReturnValue
	stat, err := os.Stat(src)
	if err != nil {
		returnvalue = model.ReturnValue{
			400,
			fmt.Sprint(err),
		}
	} else if stat.IsDir() {
		returnvalue = model.ReturnValue{
			400,
			fmt.Sprint("不能以文件夹作为拷贝源"),
		}
	}
	file, err := os.Open(src)
	if err != nil {
		returnvalue = model.ReturnValue{
			400,
			fmt.Sprint(err),
		}
	} else {
		returnvalue = model.ReturnValue{
			200,
			fmt.Sprint("ok"),
		}
	}
	defer file.Close()
	info := SerializeReturns(returnvalue)
	conn.Write(info)
	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(conn)
	defer writer.Flush()
	io.Copy(writer, reader)
}

func GetFile(conn net.Conn,dest string) {
	var returnvalue model.ReturnValue
	destfile, err := os.OpenFile(dest, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		returnvalue = model.ReturnValue{
			400,
			fmt.Sprint(err),
		}
	}
	defer destfile.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(destfile)
	defer writer.Flush()

	io.Copy(writer, reader)
	info := SerializeReturns(returnvalue)
	conn.Write(info)
}