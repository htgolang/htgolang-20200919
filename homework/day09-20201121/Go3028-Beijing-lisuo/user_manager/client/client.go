package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
)

const (
	proto   = "tcp"
	addr    = "127.0.0.1:8081"
	headLen = 5
)

// Head ...
type Head struct {
	// add, mod, show, get, del, help
	Operation string
	// id/name duplicate, phone number not a number, born time format fault
	Message string
	// noUser: 404
	// error: 500 + Message: id/name duplicate, phone number not a number, born time format fault
	// success: 200
	Status int
}

func main() {

	conn, err := net.Dial(proto, addr)
	if err != nil {
		panic(err)
	}

	h := Head{
		Operation: "help",
		Message:   "request to add a user",
		Status:    0,
	}
	WriteHead(conn, h)
	if h.Operation == "help" {
		h := ReadHead(conn)
		fmt.Println("help: ", h.Message)
	}
	conn.Close()
}

// WriteHead wrap WriteHeadLen and WriteHeadBody
func WriteHead(c net.Conn, h Head) {
	WriteHeadLen(c, h)
	WriteHeadBody(c, h)
}

// WriteHeadLen send json head len to client
func WriteHeadLen(c net.Conn, h Head) {
	bt, err := json.Marshal(h)
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
func WriteHeadBody(c net.Conn, h Head) {
	b, _ := json.Marshal(h)
	_, errW := c.Write(b)
	if errW != nil {
		c.Close()
		panic(errW)
	}
}

// ReadHead read json response head from server
func ReadHead(c net.Conn) Head {
	conLen := readHeadLen(c)
	var d = make([]byte, conLen)
	buf := bytes.NewBuffer(d)
	_, errR := c.Read(buf.Bytes())
	if errR != nil {
		c.Close()
		panic(errR)
	}
	responseBytes := buf.Bytes()
	var response = Head{}
	errUnmarshal := json.Unmarshal(responseBytes, &response)
	if errUnmarshal != nil {
		panic(errUnmarshal)
	}
	return response
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
