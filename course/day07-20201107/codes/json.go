package main

import (
	"encoding/json"
	"fmt"
)

// 结构体属性的tag tagname:"tagvalue"
// name,type,omitempty
// name,omitempty

type Addr struct {
	A string
	B string
	c string
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`

	age   int    // 首字母小写
	Tel   string `json:"tel"`
	Addr  Addr   `json:"addr,omitempty"`
	IsBoy bool   `json:"is_body"`
}

/*
Id => id => 结构体信息
Id =》 value
"id" => value
*/

func main() {
	user := User{1, "kk", "xxx", 30, "123xxx", Addr{"x", "y", "z"}, true}
	// encoding/json
	// 内存结构 => []byte (编码) 序列化
	// []byte, error = Marshal()
	// {"Id" : 1, "Name" : "kk", "Password" : "xxx", "Tel", "Add"}
	if ctx, err := json.Marshal(user); err == nil {
		fmt.Println(string(ctx))
	}
	if ctx, err := json.MarshalIndent(user, "", "\t"); err == nil {
		fmt.Println(string(ctx))
	}

	// []byte => 内存结构 (解码) 反序列化

	ctx := `{"id":100,"Name":"kk2","Password":"xxx","Tel":"123xxx"}`
	var user2 User
	fmt.Println(json.Valid([]byte(ctx)))
	if err := json.Unmarshal([]byte(ctx), &user2); err == nil {
		fmt.Printf("%#v\n", user2)
	} else {
		fmt.Println(err)
	}

}
