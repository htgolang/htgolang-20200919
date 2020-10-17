package pkg

import "fmt"

var name = ""

func init() {
	fmt.Println("pkg.init")
}

func init() {
	fmt.Println("pkg.init2")
}
