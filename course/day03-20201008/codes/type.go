package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func sayHi() {
	fmt.Println("hi")
}

func sayHello() {
	fmt.Println("hello")
}

func genFunc() func() {
	if rand.Int()%2 == 0 {
		return sayHi
	}
	return sayHello
}

func aFields(split rune) bool {
	return split == 'a'
}

func main() {
	rand.Seed(time.Now().Unix())
	f := genFunc()
	f()

	fmt.Printf("%q\n", strings.FieldsFunc("abcdefabcabc", aFields))
}
