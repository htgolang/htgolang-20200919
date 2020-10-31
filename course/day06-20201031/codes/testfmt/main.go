package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		return
	}
	defer file.Close()
	name := "kk"
	fmt.Fprintf(file, "i am: %s", name)
}
