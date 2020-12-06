// package main

// import (
// 	"fmt"
// 	"net/rpc"
// )

// type Request struct {
// 	R int
// 	L int
// }

// type Response struct {
// 	Res int
// }

// func main() {
// 	client, err := rpc.DialHTTPPath("tcp", "127.0.0.1:8888", "/")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	a := Request{10, 11}
// 	b := new(Response)
// 	err = client.Call("Action.Sum", a, b)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(b.Res)
// }
package main

import "fmt"

type Op struct {
	A int
	B int
}

func (a *Op) Sum() int {
	return a.A + a.B
}
func (c *Op) Chen() int {
	return c.A * c.B
}
func main() {
	a := Op{10, 2}
	fmt.Println("sum", a.Sum())
	fmt.Println("chen", a.Chen())
}
