package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("PATH"))
	fmt.Println(os.Getenv("RunMode"))
}
