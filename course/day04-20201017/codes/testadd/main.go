package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	usage := func() {
		fmt.Println("usage: add.exe 1 2 [3 4 ...]")
	}

	if len(args) < 2 {
		usage()
		return
	}

	total := 0

	for _, value := range args {
		if intValue, err := strconv.Atoi(value); err != nil {
			usage()
			return
		} else {
			total += intValue
		}
	}

	fmt.Println("total:", total)

}
