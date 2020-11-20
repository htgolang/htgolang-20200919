package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
		wg.Add(1)
		go fileCopy(flagSrc, flagDest, wg)
		wg.Wait()
		return
	case 1:
		fmt.Printf("%v is a dir\n", *flagSrc)
		//var walkedFilesAbs, walkedDirAbs []string
		walkedDirs, walkedFiles, err := walkReturnDirSlice(*flagSrc)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("walkedDirs: ", walkedDirs)
		fmt.Println("walkedFiles: ", walkedFiles)

		// TODO
		// there are multi dirs, not only deepest
		// should save all the dir to a []string
		// if not same dir, then os.MkdirAll

		for _, file := range walkedFiles {
			wg.Add(1)
			makedDir, err := makeDir(file, *flagDest)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("copying file: %v to: %v\n", file, makedDir)
			go fileCopy(&file, &makedDir, wg)
		}
		wg.Wait()
		return
	// not exists
	case -1:
		fmt.Printf("%v is not exist\n", *flagSrc)
		return
	default:
		fmt.Printf("strange things happened.\n")
	}
	fmt.Println(files)

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

// make dir by abs file path
func makeDir(walkedFile, destDir string) (string, error) {
	fileDir := filepath.Dir(walkedFile)
	dirToMake := filepath.Join(destDir, fileDir)
	fmt.Println("dirToMake: ", dirToMake)
	errMkdir := os.MkdirAll(dirToMake, os.ModePerm)
	if errMkdir != nil {
		return "", errMkdir
	}
	return dirToMake, nil
}

// return walked dirs slice and walked files slice
func walkReturnDirSlice(dir string) ([]string, []string, error) {
	var walkedFiles, walkedDirs []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if isDir(path) == 1 {
			walkedDirs = append(walkedDirs, path)
		} else if isDir(path) == 0 {
			walkedFiles = append(walkedFiles, path)
		}
		return nil
	})

	if err != nil {
		return []string{}, []string{}, err
	}

	return walkedDirs, walkedFiles, nil
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
	//var opt string
	srcAbs, _ := filepath.Abs(*flagSrc)
	destAbs, _ := filepath.Abs(*flagDest)
	destAbsInfo, errDestAbs := os.Stat(destAbs)
	//srcBase := filepath.Base(srcAbs)
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
