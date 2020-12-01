package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	proto   = "tcp"
	addr    = "127.0.0.1:8081"
	headLen = 5
)

// Head  represents operation and status
type Head struct {
	// add, mod, show, get, del, help
	Operation string `json:"operation"`
	// id/name duplicate, phone number not a number, born time format fault
	// user's info
	Message string `json:"message"`
	// noUser: 404
	// error: 500 + Message: id/name duplicate, phone number not a number, born time format fault
	// success: 200
	Status  int       `json:"status"`
	ID      int64     `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Cell    string    `json:"cell"`
	Born    time.Time `json:"born"`
	Passwd  string    `json:"passwd"`
}

func main() {

	conn, err := net.Dial(proto, addr)
	if err != nil {
		panic(err)
	}

	h := ParseHead(&Head{})
	WriteHead(conn, h)
	switch h.Operation {
	case "help":
		h := ReadHead(conn)
		fmt.Println("help: ", h.Message)
	case "add":
		h = InputUser(&h)
		WriteHead(conn, h)
		h = ReadHead(conn)
		fmt.Println("add return head: ", h)
		// add user {ID, Name, Address, Cell, Born, Passwd}
		// add
		// input
		// send
	case "show":
	case "mod":
	case "del":
	case "get":
	default:
		h := ReadHead(conn)
		fmt.Printf("server says: %v, Status code: %v\n", h.Message, h.Status)
		return
	}
	conn.Close()
}

// ============== cmd =============

// ParseHead gen cmd
func ParseHead(h *Head) Head {
	op := flag.String("op", "help", "specify the command")
	msg := flag.String("msg", "message", "specify message to server")
	if len(os.Args) < 2 {
		flag.Usage()
		return *h
	}
	flag.Parse()
	h.Operation = *op
	h.Message = *msg
	return *h
}

func InputUser(h *Head) Head {
	fmt.Printf("ID: \n> ")
	id, err := strconv.ParseInt(Read(), 10, 64)
	if err != nil {
		panic(err)
	}
	h.ID = id
	fmt.Printf("Name: \n> ")
	h.Name = Read()
	fmt.Printf("Address: \n> ")
	h.Address = Read()
	fmt.Printf("Cell: \n> ")
	h.Cell = Read()
	fmt.Printf("Born: \n> ")
	str := Read()
	t, err := time.Parse("2006.01.02", str)
	if err != nil {
		panic(err)
	}
	h.Born = t
	fmt.Printf("Passwd: \n> ")
	h.Passwd = Read()
	return *h
}

// ============== protocol =============

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

// Read read content from standard input
func Read() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	line := strings.TrimSpace(scanner.Text())
	return line
}
