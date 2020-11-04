package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func allGoFiles() (ret *[]string) {
	gofiles := make([]string, 0)
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return nil
		}
		if !info.IsDir() {
			if filepath.Ext(filepath.Base(path)) == ".go" || filepath.Ext(filepath.Base(path)) == ".cgo" {
				gofiles = append(gofiles, path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path: %v\n", err)
		return nil
	}
	return &gofiles
}

func readLine(filename string) (lineSum int, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}

	reader := bufio.NewReader(file)
	for {
		_, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}
		if !isPrefix {
			lineSum++
		}
	}
	return
}
