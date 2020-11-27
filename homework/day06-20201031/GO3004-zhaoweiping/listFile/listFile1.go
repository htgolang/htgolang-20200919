package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	var srcFile string
	flag.StringVar(&srcFile, "s", "", "源文件")
	flag.Parse()

	files, _ := ioutil.ReadDir(srcFile)
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
