package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"zhao/pathtype"
)

func main() {
	n, _ := os.Executable()
	cmdname := filepath.Base(n)
	help := flag.Bool("h", false, "help message")
	// force := flag.Bool("f", false, "force overwrite")

	flag.Usage = func() {
		fmt.Printf(`
Usage %s [-h] [-f] source dest
Options: 

		`, cmdname)
		flag.PrintDefaults()
	}
	flag.Parse()

	if *help || len(flag.Args()) != 2 {
		flag.Usage()
		return
	}

	source := flag.Args()[0]
	dest := flag.Args()[1]

	s, _ := filepath.Abs(source)
	d, _ := filepath.Abs(dest)
	pathtype.Run(s, d)
}
