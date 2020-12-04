package userutils

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
)

type Cmd struct {
	Cmd string
	Code int
	Data string
	Userinfo map[string]string
}

type Msg struct {
	Status int
	Ack bool
	Data string
}

func Readconn(conn net.Conn) (Cmd,error) {
	tmplen := make([]byte, 5)
	_, err := conn.Read(tmplen)
	if err != nil {
		return Cmd{},err
	}
	length, err := strconv.Atoi(string(tmplen))
	if err != nil {
		return Cmd{},err
	}
	data := make([]byte, length)
	_, err = conn.Read(data)
	if err != nil {
		return Cmd{},err
	}
	var cmd Cmd
	err = json.Unmarshal(data, &cmd)
	if err != nil {
		fmt.Printf("json unmarshal error:%s", err)
		return Cmd{}, err
	}
	return cmd, nil
}

func Sendmes(conn net.Conn, dmsg Msg) error  {
	bytedata, err := json.Marshal(dmsg)
	if err != nil {
		fmt.Println(err)
		return  err
	}
	length := len(bytedata)
	_, err = conn.Write([]byte(fmt.Sprintf("%05d", length)))
	if err != nil {
		return err
	}
	_, err = conn.Write(bytedata)
	if err != nil {
		return err
	}
	return nil

}

func SendCmd(conn net.Conn, dcmd Cmd) error  {
	bytedata, err := json.Marshal(dcmd)
	if err != nil {
		fmt.Println(err)
		return  err
	}
	length := len(bytedata)
	_, err = conn.Write([]byte(fmt.Sprintf("%05d", length)))
	if err != nil {
		return err
	}
	_, err = conn.Write(bytedata)
	if err != nil {
		return err
	}
	return nil

}

func ReadMsg(conn net.Conn) (Msg,error) {
	tmplen := make([]byte, 5)
	_, err := conn.Read(tmplen)
	if err != nil {
		return Msg{},err
	}
	length, err := strconv.Atoi(string(tmplen))
	if err != nil {
		return Msg{},err
	}
	data := make([]byte, length)
	_, err = conn.Read(data)
	if err != nil {
		return Msg{},err
	}
	msg := new(Msg)
	err = json.Unmarshal(data, msg)
	if err != nil {
		fmt.Printf("json unmarshal error:%s", err)
		return Msg{}, err
	}
	return *msg, nil
}