package main

import "fmt"

type IPAndStatusCode struct {
	ip   string
	code int
}

func main() {
	v1 := IPAndStatusCode{"1.1.1.1", 200}
	v2 := IPAndStatusCode{"1.1.1.1", 200}
	v3 := IPAndStatusCode{"1.1.1.2", 200}
	v4 := IPAndStatusCode{"1.1.1.1", 201}

	fmt.Println(v1 == v2)
	fmt.Println(v1 == v3)
	fmt.Println(v1 == v4)
}
