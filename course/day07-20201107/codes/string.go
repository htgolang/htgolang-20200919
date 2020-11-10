package main

import (
	"fmt"
)

type User struct {
	Id       int
	Name     string
	Password string
}

func (user User) String() string {
	return fmt.Sprintf("User[Name=%s]", user.Name)
}

func main() {
	var user User = User{0, "xxx", "yyy"}
	fmt.Println(user)
	puser := &User{1, "xxx", "yyy"}
	fmt.Println(puser)

}
