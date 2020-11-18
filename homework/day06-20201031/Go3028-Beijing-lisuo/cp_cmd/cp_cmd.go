package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var destBase, destDir string

var destFile *os.File

// Abs = Dir/Base

func main() {
	// cmd flag define
	flagSrc := flag.String("src", "/bar/foo", "specify source file location")
	flagDest := flag.String("dest", "/foo/bar", "sepcify dest file location")
	flag.Usage = func() {
		fmt.Println("Usage: go run cp_cmd.go --src ./bar/foo --dest /foo/bar")
	}
	// parse flag
	flag.Parse()

	srcAbs, _ := filepath.Abs(*flagSrc)
	destAbs, _ := filepath.Abs(*flagDest)
	destAbsInfo, errDestAbs := os.Stat(destAbs)
	srcBase := filepath.Base(srcAbs)
	destBase = filepath.Base(destAbs)
	destDir = filepath.Dir(destAbs)

	// if user not specify a file in src, prompt error
	srcAbsInfo, err := os.Stat(srcAbs)
	if os.IsNotExist(err) {
		log.Fatal(err)
		return
	} else if srcAbsInfo.IsDir() {
		fmt.Printf("the file \"%v\" is a dir, please specify a file.\n",
			srcBase)
		return
	}

	// copy
	joinSrcDirToDestBase(destAbsInfo, errDestAbs, &srcAbs, &destAbs)
	fileCopy(&srcAbs, &destAbs, &srcBase, &destDir)

}

// if the dest not contains the file name, then join
// the file name in src to dest dir
func joinSrcDirToDestBase(fi os.FileInfo, e error, srcAbs, destAbs *string) {
	// if user not specify the file name in dest, make it happen
	if e != nil {
		log.Fatal(e)
	} else if fi.IsDir() {
		(*destAbs) = filepath.Join((*destAbs), filepath.Base(*srcAbs))
		destDir = filepath.Dir(*destAbs)
		destBase = filepath.Base(*destAbs)
	}

}

// file copy op
func fileCopy(srcAbs, destAbs, srcBase, destDir *string) {
	var opt string

	// file copy op
	srcFile, err := os.Open(*srcAbs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("srcFile: ", srcFile)
	defer srcFile.Close()

	srcReader := bufio.NewReader(srcFile)
	//fmt.Println("srcReader: ", srcReader)

	// if dest dir already have a same name file, abort
	fileInfoList, err := ioutil.ReadDir(*destDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range fileInfoList {
		if file.Name() == *srcBase {
			fmt.Printf("There's also a file named %v in %v\nAre you want to overwrite it?(y/n): ", *srcBase, *destDir)
			fmt.Scanln(&opt)
			opt = strings.TrimSpace(strings.ToLower(opt))
			if opt == "y" {
				destFile, err := os.Create(*destAbs)
				fmt.Printf("destFile type: %T, value: %v: \n", destFile, destFile)
				defer destFile.Close()
				if err != nil {
					log.Fatal(err)
				}
				srcReader.WriteTo(destFile)
				return
			} else if opt == "n" {
				fmt.Println("Nothing changed.")
				return
			} else {
				fmt.Println("Nothing changed.")
				return
			}

		}
	}
	destFile, err := os.Create(*destAbs)
	fmt.Printf("destFile type: %T, value: %v: \n", destFile, destFile)
	defer destFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, errww := srcReader.WriteTo(destFile)
	if errww != nil {
		log.Fatal(errww)
	}

}
