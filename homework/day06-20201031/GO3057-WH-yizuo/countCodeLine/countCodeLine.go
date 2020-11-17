package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var fileList []string

// Print all files in the directory
func findFile(str string)  {
	stat, err := os.Lstat(str)
	// Check for presence
	if os.IsNotExist(err) {
		fmt.Printf("Not %v doesn't exist.\n",str)
		return
	}
    // Check is not a directory
	if !stat.IsDir() {
		fmt.Printf("The %v is not a directory.\n",str)
		return
	}

    // Print all files
	filepath.Walk(str, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir(){
			if strings.HasSuffix(info.Name(), ".go") || strings.HasSuffix(info.Name(), ".cgo") {
				fileList = append(fileList,path)
				return err
			}
			return err
		}
		return err
	})
}

// Calculates the number of lines of code for the file based on the absolute path of the file entered
func countFileLine(fileName string)(num int ){
	file,err := os.Open(fileName)
	if err != nil {
		fmt.Println(fileName)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for  {
		_,isPrefix,err := reader.ReadLine()
		if err != nil {
			return
		}
		if !isPrefix {
			num++
		}
	}
	return num
}


// main
func main()  {
    // run
	if len(os.Args) < 2 {
    	fmt.Println("Please enter the directory after the script.")
	} else {
		// Get file list
		findFile(os.Args[1])
		// Number of lines of code in the print file list
		for _,v := range fileList {
			fmt.Printf("脚本路径:%v\t代码行数:%v\n",v,countFileLine(v))
		}
	}
}