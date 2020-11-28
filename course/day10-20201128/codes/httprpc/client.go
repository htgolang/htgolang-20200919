package main

import (
	"fmt"
	"net/rpc"
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

func main() {
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:8888")

	req := AddRequest{3, 10}
	resp := AddResponse{}

	err := client.Call("Calc.Add", req, &resp)
	fmt.Println(err, resp)

	client.Close()
}
