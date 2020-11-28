package main

import (
	"fmt"
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

func main() {
	client, _ := jsonrpc.Dial("tcp", "127.0.0.1:8888")
	req := AddRequest{1, 100}
	resp := AddResponse{}

	// err := client.Call("AAAA.Add", req, &resp)
	// fmt.Println(err, resp)
	// 同步调用，异步调用

	call := client.Go("AAAA.Add", req, &resp, nil)
	for {
		select {
		case result := <-call.Done:
			fmt.Println(result.Reply, result.Error)
		default:
			fmt.Println(time.Now())
			time.Sleep(2 * time.Second)
		}

	}
}
