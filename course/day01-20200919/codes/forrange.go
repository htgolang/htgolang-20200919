package main

import "fmt"

func main() {
	s := "我爱中国"
	for i, v := range s {
		fmt.Printf("%d, %q\n", i, v)
	}
}
