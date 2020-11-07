package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 结构体属性的tag tagname:"tagvalue"
// name,type,omitempty
// name,omitempty
type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`

	age   int    // 首字母小写
	Tel   string `json:"tel"`
	Addr  string `json:"addr,omitempty"`
	IsBoy bool   `json:"is_body"`
}

func main() {
	// users := []User{
	// 	{1, "kk", "xxx", 30, "123xxx", "xxx", true},
	// 	{2, "kk2", "xxx", 30, "123xxx", "xxx", true},
	// 	{3, "kk3", "xxx", 30, "123xxx", "xxx", true},
	// }

	// encoder
	// file, _ := os.Create("users.json")
	// defer file.Close()

	// encoder := json.NewEncoder(file)
	// encoder.Encode(users)

	// decoder
	file, _ := os.Open("users.json")
	defer file.Close()
	decoder := json.NewDecoder(file)

	var users2 []User
	decoder.Decode(&users2)

	fmt.Println(users2)
}
