package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/hpcloud/tail"
)

var (
	filepath string
	num      int
)

func getOffset(filepath string, num int) (ret int64) {
	fileHandle, err := os.Open(filepath)

	if err != nil {
		panic("Cannot open file")
	}
	defer fileHandle.Close()

	ret = 0
	linecount := 0
	stat, _ := fileHandle.Stat()
	filesize := stat.Size()
	for {
		ret--
		fileHandle.Seek(ret, io.SeekEnd)

		char := make([]byte, 1)
		fileHandle.Read(char)

		if ret != -1 && (char[0] == 10 || char[0] == 13) { // stop if we find a line
			linecount++
			continue
		}

		if ret == -filesize { // stop if we are at the begining
			break
		}

		if linecount >= num {
			break
		}
	}

	ret += 2

	return
}

//Tail ...
func Tail() {
	config := tail.Config{
		ReOpen:    true,                                                        // 重新打开
		Follow:    true,                                                        // 是否跟随
		Location:  &tail.SeekInfo{Offset: getOffset(filepath, num), Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                                       // 文件不存在不报错
		Poll:      true,
	}
	tails, err := tail.TailFile(filepath, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:", line.Text)
	}
}

func main() {
	flag.IntVar(&num, "n", 5, "number of flush line, default is 5")
	flag.StringVar(&filepath, "f", "", "file to tail")
	flag.Usage = func() {
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "Note: easy version of tail\n")
	}
	flag.Parse()
	if filepath == "" {
		fmt.Println("Please give a filename.")
		return
	}
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) || info.IsDir() {
		fmt.Println("File does not exists or not a regular file.")
		return
	}
	Tail()
}
