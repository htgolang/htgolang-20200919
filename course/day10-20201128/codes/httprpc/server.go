package main

import (
	"fmt"
	"net/http"
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

type Calc struct{}

// 参数1：请求对象(可以是指针/值)
// 参数2：响应对象(可以是指针)
// 返回值: error
func (c *Calc) Add(req AddRequest, resp *AddResponse) error {
	fmt.Println("calc.add")
	resp.Result = req.Left + req.Right
	return nil
}

func main() {
	rpc.Register(&Calc{})

	rpc.HandleHTTP() // 使用http请求
	http.ListenAndServe("0.0.0.0:8888", nil)
}
