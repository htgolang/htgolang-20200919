package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

// 远程服务 Add(1, 2) int
type AddRequest struct {
	Left  int
	Right int
}

// type AddResponse int
type AddResponse struct {
	Result int
}

type Calc struct{}

// 参数1：请求对象(可以是指针/值)
// 参数2：响应对象(可以是指针)
// 返回值: error
func (c *Calc) Add(req AddRequest, resp *AddResponse) error {
	fmt.Println("calc.add")
	time.Sleep(10 * time.Second)
	resp.Result = req.Left + req.Right
	return nil
}

func main() {
	rpc.Register(&Calc{})
	rpc.RegisterName("AAAA", &Calc{})

	listener, _ := net.Listen("tcp", "0.0.0.0:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		go jsonrpc.ServeConn(conn)
	}
	listener.Close()
}
