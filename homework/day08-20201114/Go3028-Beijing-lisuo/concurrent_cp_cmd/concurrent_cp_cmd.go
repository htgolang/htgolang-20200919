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
	"sync"
)

var (
	destBase, destDir string
	destFile          *os.File
	files             []string
	copyWorkers       int
)

func main() {
	//To implement
	// cmd flag define
	flagSrc := flag.String("src", "/bar/foo", "specify source file location")
	flagDest := flag.String("dest", "/foo/bar", "sepcify dest file location")
	flag.Usage = func() {
		fmt.Println("Usage: go run cp_cmd.go --src ./bar/foo --dest /foo/bar")
	}
	// parse flag
	flag.Parse()
	wg := &sync.WaitGroup{}

	n := isDir(*flagSrc)
	switch n {
	// file
	case 0:
		fmt.Printf("%v is a file\n", *flagSrc)
		fileCopy(flagSrc, flagDest, wg)
		return
	// dir
	case 1:
		fmt.Printf("%v is a dir\n", *flagSrc)
		deepestFile, err := walkReturnDeepestDir(*flagSrc)
		if err != nil {
			fmt.Println(err)
			return
		}
		// TODO
		// there are multi dirs, not only deepest
		// should save all the dir to a []string
		// if not same dir, then os.MkdirAll
		errMkdir := os.MkdirAll(filepath.Dir(deepestFile), os.ModePerm)
		if errMkdir != nil {
			fmt.Println(errMkdir)
			return
		}
		files, err := listFiles(*flagSrc)
		if err != nil {
			fmt.Println(err)
			return
		}
		// TODO
		// save all walked file to a []string, then os.MkdirAll
		// then calc the file amount, start corresponding goroutines to cp
		// --src /path/to/file --desc /path/to/file
		return
	// not exists
	case -1:
		fmt.Printf("%v is not exist\n", *flagSrc)
		return
	default:
		fmt.Printf("strange things happened.\n")
	}
	wg.Add(1)
	go fileCopy(flagSrc, flagDest, wg)
	wg.Wait()
	fmt.Println(files)
	//files, err := listFiles("../concurrent_cp_cmd")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("files: ", files)
}

// dir return 1, file return 0, not exist return -1
func isDir(file string) int {
	f, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return -1
		}
	}
	if f.IsDir() {
		return 1
	}
	return 0
}

func walkReturnDeepestDir(dir string) (string, error) {
	var walkedFiles []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		walkedFiles = append(walkedFiles, path)
		return nil
	})

	if longest := allLongestStrings(walkedFiles); len(longest) != 0 {
		return longest[0], nil
	}

	if err != nil {
		return "", err
	}
}

// list files in a dir, return as a []string
func listFiles(dir string) ([]string, error) {
	dirInfo, err := os.Stat(dir)
	if err != nil {
		fmt.Println(err)
		return files, err
	}
	if os.IsNotExist(err) {
		fmt.Println(err)
	}
	if dirInfo.IsDir() {
		filesInfo, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Println(err)
			return files, err
		}
		for _, fInfo := range filesInfo {
			files = append(files, fInfo.Name())
		}

	}
	return files, nil
}

// file copy op
func fileCopy(flagSrc, flagDest *string, wg *sync.WaitGroup) {
	var opt string
	srcAbs, _ := filepath.Abs(*flagSrc)
	destAbs, _ := filepath.Abs(*flagDest)
	destAbsInfo, errDestAbs := os.Stat(destAbs)
	srcBase := filepath.Base(srcAbs)
	destBase = filepath.Base(destAbs)
	destDir = filepath.Dir(destAbs)

	// if user not specify a file in src, prompt error
	_, err := os.Stat(srcAbs)
	if os.IsNotExist(err) {
		log.Fatal(err)
		return
	}
	// if the dest not contains the file name, then join
	// the file name in src to dest dir
	func(fi os.FileInfo, e error, srcAbs, destAbs *string) {
		// if user not specify the file name in dest, make it happen
		if e != nil {
			log.Fatal(e)
		} else if fi.IsDir() {
			(*destAbs) = filepath.Join((*destAbs), filepath.Base(*srcAbs))
			destDir = filepath.Dir(*destAbs)
			destBase = filepath.Base(*destAbs)
		}
	}(destAbsInfo, errDestAbs, &srcAbs, &destAbs)

	// file copy op
	srcFile, err := os.Open(srcAbs)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("srcFile: ", srcFile)
	defer srcFile.Close()

	srcReader := bufio.NewReader(srcFile)

	// if dest dir already have a same name file, abort
	fileInfoList, err := ioutil.ReadDir(destDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range fileInfoList {
		if file.Name() == srcBase {
			fmt.Printf("There's also a file named %v in %v\nAre you want to overwrite it?(y/n): ", srcBase, destDir)
			fmt.Scanln(&opt)
			opt = strings.TrimSpace(strings.ToLower(opt))
			if opt == "y" {
				destFile, err := os.Create(destAbs)
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
	destFile, err := os.Create(destAbs)
	fmt.Printf("destFile type: %T, value: %v: \n", destFile, destFile)
	defer destFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, errww := srcReader.WriteTo(destFile)
	if errww != nil {
		log.Fatal(errww)
	}
	wg.Done()

}

func allLongestStrings(inputArray []string) []string {
	n := 0
	max := -1
	for i, s := range inputArray {
		if len(s) < max {
			// Skip shorter string
			continue
		}
		if len(s) > max {
			// Found longer string. Update max and reset result.
			max = len(s)
			n = 0
		}
		inputArray[n], inputArray[i] = inputArray[i], inputArray[n]
		n++
	}
	return inputArray[:n]
}
