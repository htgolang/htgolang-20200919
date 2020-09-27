package main

import (
	"bufio"
	"fmt"
	"os"
)

//CharStat ...
func CharStat() {
	file, err := os.Open("./i-have-a-dream.txt")
	if err != nil {
		fmt.Printf("error when open file, %v\n", err)
		return
	}
	defer file.Close()

	CharMap := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0, "n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, v := range line {
			_, ok := CharMap[string(v)]
			if !ok {
				continue
			} else {
				CharMap[string(v)]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error for scanner, %v\n", err)
		return
	}

	for k, v := range CharMap {
		fmt.Printf("char: %v --- count: %v\n", k, v)
	}
}
