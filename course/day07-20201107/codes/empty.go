package main

import "fmt"

type Empty interface {
}

type User struct {
	Name     string
	Password string
}

func main() {
	var empty Empty
	empty = 1
	fmt.Printf("%T, %#v\n", empty, empty)

	if v, ok := empty.(int); ok {
		fmt.Println(v + 2)
	}

	empty = "xxxxx"
	fmt.Printf("%T, %#v\n", empty, empty)

	// empty + "xxx"

	empty = true
	fmt.Printf("%T, %#v\n", empty, empty)

	// empty && true

	empty = User{"x", "y"}

	if u, ok := empty.(User); ok {
		fmt.Println(u.Name, u.Password)
	}

	// empty.Name
	fmt.Printf("%T, %#v\n", empty, empty)
}
