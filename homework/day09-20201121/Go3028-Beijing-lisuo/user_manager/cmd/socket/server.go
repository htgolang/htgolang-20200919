package socket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day09-20201121/Go3028-Beijing-lisuo/user_manager/cmd/funcs"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day09-20201121/Go3028-Beijing-lisuo/user_manager/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day09-20201121/Go3028-Beijing-lisuo/user_manager/utils"
)

const (
	proto   = "tcp"
	addr    = ":8081"
	headLen = 5
)

var helpMsg = `
+-------+---------------------+
|  CMD  |      Function       |
+-------+---------------------+
| help  | ShowHelp            |
| add   | AddUser             |
| show  | ShowCurrentUserList |
| mod   | ModifyUser          |
| del   | DelUser             |
| get   | QueryUser           |
+-------+---------------------+
`

// Head  represents operation and status
type Head struct {
	Operation string    `json:"operation"`
	Message   string    `json:"message"`
	Status    int       `json:"status"`
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Cell      string    `json:"cell"`
	Born      time.Time `json:"born"`
	Passwd    string    `json:"passwd"`
}

/*

+-------+---------------------+
|  CMD  |      Function       |
+-------+---------------------+
| get   | QueryUser           |
| h     | ShowHelp            |
| show  | ShowCurrentUserList |
| q     | utils.Quit          |
| del   | DelUser             |
| help  | ShowHelp            |
| cls   | utils.ClearScreen   |
| quit  | utils.Quit          |
| mycmd | DoMap               |
| rot   | Rotate              |
| add   | AddUser             |
| mod   | ModifyUser          |
| save  | SaveUsers           |
| Q     | utils.Quit          |
| exit  | utils.Quit          |
+-------+---------------------+

*/

// Server for remote user manager
func Server() {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		panic(err)
	}
	conn, errA := listener.Accept()
	if errA != nil {
		panic(errA)
	}

	res := ReadHead(conn)
	fmt.Println("Head: ", res)

	switch res.Operation {
	case "help":
		res.Message = helpMsg
		res.Status = 200
		showClientHelp(conn, &res)
	case "add":
		res = ReadHead(conn)
		HandleAddUser(conn, &res)
		fmt.Printf("user: %v added.\n", res.Name)
	case "show":
	case "mod":
	case "del":
	case "get":
	default:
		fmt.Println("something strange happened")
		res.Message = "something strange happened"
		res.Status = 500
		WriteHead(conn, res)
	}

	conn.Close()

}

func showClientHelp(c net.Conn, h *Head) {
	WriteHead(c, *h)
}

func HandleAddUser(c net.Conn, h *Head) {
	addUser(c, h)
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

// ============ user op ===============
func addUser(c net.Conn, h *Head) {
	ul := &define.UserList
	var uc define.User
	var Name string
	ID := funcs.GetMaxID(ul) + 1
	Name = h.Name
	user, err := funcs.NameFindUser(ul, Name)
	// find the user name, so prompt reinput
	if err == nil {
		h.Message = "The person already exists: " + user.Name
		WriteHead(c, *h)
	} else if Name == "" {
		h.Message = "damn fool, must specify a Name"
		WriteHead(c, *h)
	}
	Cell := h.Cell
	if !utils.JustDigits(Cell) {
		h.Message = "Please input a real cell number"
		WriteHead(c, *h)
	}
	Address := h.Address
	Born := h.Born.Format("2006.01.02")
	if err := utils.DateCheck(Born); err != nil {
		h.Message = "Please input a legal born time[YYYY.MM.DD]: "
		WriteHead(c, *h)
	}
	Passwd := h.Passwd
	uc = funcs.NewUser(ID, Name, Cell, Address, Born, Passwd)
	define.UserList = append(*ul, uc)
	h.Status = 200
	h.Message = "user: " + h.Name + " added."
	WriteHead(c, *h)
}
