package main

import "fmt"

type A struct {
	Name string
}

func (b *A) Test() {
	// a = &A{"xxxxx"}
	*b = A{"xxxxxx"}
	b = &A{"xxxxxxx"}
	//
	b.Name = "xxx"
}


func main() {
	a := &A{"aaaaaaaa"}
	b := a
	*b = A{"xxxxxxx"}

	b = &{"ccccccccc"}
	a.Test()
	fmt.Print(a.Name)
}
