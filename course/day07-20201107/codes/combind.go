package main

import "fmt"

// 发送接口
type Sender interface {
	Send(string) error
}

// 代码的复用
// 网络连接 抽象(接口)
// Open, Send, Recive, Close
type Connection interface {
	Sender // 匿名组合Sender接口

	Open() error
	// Send(string) error
	Recive() (string, error)
	Close() error
}

type TcpConnection struct {
}

func (c *TcpConnection) Open() error {
	return nil
}

func (c *TcpConnection) Send(msg string) error {
	return nil
}

func (c *TcpConnection) Recive() (string, error) {
	return "", nil
}

func (c *TcpConnection) Close() error {
	return nil
}

// 结构体组合接口
type Client struct {
	Connection
	C    Connection
	Name string
}

func main() {
	var conn Connection = new(TcpConnection)
	fmt.Printf("%T, %#v\n", conn, conn)
	conn.Open()
	conn.Send("xxx")
	conn.Recive()
	conn.Close()

	client := Client{new(TcpConnection), new(TcpConnection) "KK"}
	client.Open()
	client.Send("xxx")
	client.Recive()
	client.Close()

	client.C.Open()
	client.C.Send("xxxx")
	client.C.Recive()
	client.C.Close()
}
