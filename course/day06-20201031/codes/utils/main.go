package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	return
	// }
	// defer file.Close()

	// fmt.Println(ioutil.ReadAll(file))
	// fmt.Println(ioutil.ReadFile("test.txt"))
	// fmt.Println(ioutil.ReadDir("a"))
	// ioutil.WriteFile("test.txt", []byte("xx123123xxx"), os.ModePerm)
	file, err := os.Open("test.txt")
	fmt.Println(err)
	defer file.Close()

	io.Copy(os.Stdout, file)
}
