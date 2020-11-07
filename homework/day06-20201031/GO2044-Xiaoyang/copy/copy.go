package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func FileIsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
		return false
	}
}

func Pflag() (string, string) {
	var (
		sFile string
		dFile string
	)
	flag.StringVar(&sFile, "s", "", "Please input srcFilename..")
	flag.StringVar(&dFile, "d", "", "Please input dstFilename..")

	flag.Parse()

	flag.Usage = func() {
		flag.PrintDefaults()
	}

	if sFile == "" && dFile == "" {
		flag.Usage()
		os.Exit(0)
	}
	return sFile, dFile
}

func Input() string {

	var (
		info string
	)
	fmt.Printf("The target file already exists,Whether or not to continue (y|n):  ")
	fmt.Scan(&info)
	return info

}

func srcFile(sfilename string) string {
	sFile, err := os.Open(sfilename)
	if err != nil {
		log.Fatal(err)
	}
	defer sFile.Close()

	ctx := make([]byte, 1024)
	for {
		n, err := sFile.Read(ctx)
		if err != err {
			log.Fatal(err)
		}
		return string(ctx[:n])
	}
}

func writeFile(sfile, dfilename string) {
	dFile, err := os.Create(dfilename)
	if err != nil {
		log.Fatal(err)
	}
	defer dFile.Close()
	dFile.WriteString(sfile)
}

func destFile(sfile, dfilename string, status bool) {
	if status != true {
		writeFile(sfile, dfilename)
	} else {
		info := Input()
		if info == "y" || info == "Y" {
			writeFile(sfile, dfilename)
		} else {
			fmt.Println("Program exits")
		}

	}

}

func main() {
	sfilename, dfilename := Pflag()
	sfile_info := srcFile(sfilename)
	status := FileIsExists(dfilename)
	destFile(sfile_info, dfilename, status)
}
