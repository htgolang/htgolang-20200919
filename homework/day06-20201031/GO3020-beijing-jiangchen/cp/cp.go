package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var srcFile = ""
var destFile = ""

func usage() {
	fmt.Fprintf(os.Stderr, "cp -s srcFile -d dstFile\n")
	flag.PrintDefaults()
}

func main() {
	flag.StringVar(&srcFile, "s", "", "source file path")
	flag.StringVar(&destFile, "d", "", "destination file path")
	flag.Usage = usage
	flag.Parse()

	if srcFile == "" || destFile == "" {
		fmt.Println("Please provide valid srcFile path and valid destFile path...")
		return
	}

	fi, err := os.Stat(srcFile)
	if os.IsNotExist(err) {
		fmt.Println("source file does not exists...")
		return
	} else if fi.Mode().IsDir() {
		fmt.Println("source is directory, but not file...")
		return
	}

	input, err := ioutil.ReadFile(srcFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	var check string

	if _, err = os.Stat(destFile); !os.IsNotExist(err) {
		for {
			fmt.Print("destination file exists, do you want to overwrite? (y/n)")
			fmt.Scanln(&check)
			if check != "y" && check != "n" {
				continue
			} else if check == "n" {
				fmt.Println("abort...")
				return
			} else {
				break
			}
		}
	}

	err = ioutil.WriteFile(destFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destFile)
		fmt.Println(err)
		return
	}
}
