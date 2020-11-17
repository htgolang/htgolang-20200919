package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)


// Print all files in the directory
func printFile(str string)  {
	stat, err := os.Lstat(str)
	// Check for presence
	if os.IsNotExist(err){
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
				fmt.Println(path)
				return err
			}
			return err
		}
		return err
	})
}


func main()  {
    // run
    if len(os.Args) < 2 {
    	fmt.Println("Please enter the directory after the script.")
	} else {
		printFile(os.Args[1])
	}
}