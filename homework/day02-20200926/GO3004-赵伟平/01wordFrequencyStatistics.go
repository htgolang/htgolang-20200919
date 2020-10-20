package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

func main() {
	inputFile := "dream.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File error: %s\n", err)
	}

	count := make(map[rune]int)
	buf1 := string(buf)

	for _, v := range strings.ToLower(buf1) {
		isLower := unicode.IsLower(v)
		if isLower {
			count[v]++
		}
	}

	// fmt.Println(count)
	for i, v := range count {
		fmt.Printf("%c: %v\n", i, v)
	}
}
