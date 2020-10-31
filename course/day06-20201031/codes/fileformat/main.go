package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type User struct {
	ID   int
	Name string
}

func main() {
	// json(文本), csv(文本), gob(二进制, go特有的，不能跨语言)
	// users := []User{
	// 	User{1, "kk"},
	// 	User{2, "libin"},
	// }

	gob.Register(User{})
	// 解码
	file, err := os.Open("users.gob")
	if err != nil {
		return
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)

	var users User
	fmt.Println(decoder.Decode(&users))
	fmt.Println(users)
}
