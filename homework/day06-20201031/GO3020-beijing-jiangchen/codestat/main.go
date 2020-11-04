package main

import (
	"fmt"
)

func main() {
	gofiles := allGoFiles()
	for _, file := range *gofiles {
		lineSum, err := readLine(file)
		if err != nil {
			fmt.Println("readLine error: ", err)
			return
		}
		fmt.Printf("filename: %s, lineSum: %d\n", file, lineSum)
	}
}
