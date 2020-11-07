package main

import "fmt"

type User struct {
	Name     string
	Password string
}

func main() {
	var empty interface{}
	empty = 1
	empty = "xxxx"
	empty = User{"x", "y"}

	if u, ok := empty.(User); ok {
		fmt.Println(u.Name, u.Password)
	}

	// empty.Name
	fmt.Printf("%T, %#v\n", empty, empty)
	// io.Copy()
}
