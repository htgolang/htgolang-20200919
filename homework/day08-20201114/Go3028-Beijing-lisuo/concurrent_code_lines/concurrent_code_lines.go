package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	//To implement
	flagSrc := flag.String("codeFolder", "/bar/foo", "specify source file location")
	flag.Usage = func() {
		fmt.Println("Usage: go run concurrent_code_lines.go --codeFolder ./bar/foo")
	}
	// parse flag
	flag.Parse()
	wg := &sync.WaitGroup{}

	n := isDir(*flagSrc)

	switch n {
	// file
	case 0:
		var linesRec = make(chan int, 1)
		wg.Add(1)
		calcLines(*flagSrc, linesRec, wg)
		wg.Wait()

	// folder
	case 1:
		fmt.Printf("%v is a dir\n", *flagSrc)
		_, walkedFiles, err := walkReturnDirSlice(*flagSrc)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("walkedFiles: ", walkedFiles)

		// buffered chan linesRec save each .go's lines
		var linesRec = make(chan int, len(walkedFiles))
		var totalLines int
		for _, file := range walkedFiles {
			wg.Add(1)
			go calcLines(file, linesRec, wg)
		}
		wg.Wait()

		// range the filled chan, calculate total lines
		for i := 0; i < len(linesRec); i++ {
			if v, ok := <-linesRec; ok {
				totalLines += v
			} else {
				close(linesRec)
			}
		}
		fmt.Printf("all lines are: %v\n", totalLines)
		return
	// not exists
	case -1:
		fmt.Printf("%v is not exist\n", *flagSrc)
		return
	default:
		fmt.Printf("strange things happened.\n")
	}
}

// calcLines calculate a single .go file's lines
func calcLines(goFile string, linesRec chan int, wg *sync.WaitGroup) {
	var lines int
	absFile, err := filepath.Abs(goFile)
	if err != nil {
		fmt.Println(err)
		wg.Done()
	}
	f, err := os.Open(absFile)
	if err != nil {
		fmt.Println(err)
		wg.Done()
	}
	defer f.Close()
	suffix := filepath.Ext(absFile)

	if suffix == ".go" {
		scanner := bufio.NewScanner(f)
		for {
			// for each scan, inc 1 of lines
			if scanner.Scan() {
				scanner.Text()
				(lines)++
			} else {
				break
			}
		}
		linesRec <- lines
		fmt.Printf("file %v have %v lines\n", goFile, lines)
	}
	wg.Done()
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
