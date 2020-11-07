package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	path := "test"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileInfos, err := file.Readdir(-1)

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() == true {
			file2, err := os.Open(path + "/" + fileInfo.Name())
			if err != nil {
				log.Fatal(err)
			}
			defer file2.Close()
			fileInfos2, err := file2.Readdir(-1)
			for _, fileinfo2 := range fileInfos2 {
				fmt.Println(filepath.Join(path, fileInfo.Name(), fileinfo2.Name()))
			}
		} else {
			fmt.Println(filepath.Join(path, fileInfo.Name()))

		}

	}
}
