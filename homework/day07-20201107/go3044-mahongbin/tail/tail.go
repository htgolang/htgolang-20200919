package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func allLineCount(path string) int {
	content, _ := ioutil.ReadFile(path) //content string
	return strings.Count(string(content), "\n") + 1
}

func fileCheck(path string) bool {
	stat, err := os.Stat(path)
	if stat.IsDir() {
		fmt.Println(path, "Is DIR")
		return false
	}
	return os.IsExist(err)
}

func tail(lastNLines int, path string, follow bool) {
	// fmt.Println("tail...", lastNLines, follow, path) // debug
	if !fileCheck(path) {
		return
	}
	fp, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件时出错", err.Error())
		return
	}
	defer fp.Close()

	countLines := allLineCount(path)
	lineNum := 0

	/* NewReader 函数用来返回一个默认大小 buffer 的 Reader 对象（默认大小是 4096） 等同于 NewReaderSize(rd,4096)
	 func NewReader(rd io.Reader) *Reader
	初始化一个Reader对象，存在buffer中的，读取一次就会被清空
	*/
	rdr := bufio.NewReader(fp)

	for {
		lineNum++
		line, err := rdr.ReadString('\n')

		//优先判断err的各种情况
		switch err {
		case io.EOF:
			// fmt.Printf("EOF -- Lines: %d", lineNum) // debug
			fmt.Printf(line) //打印EOF所在行，没有\n
			// 不加 -f 则退出循环
			if !follow {
				return
			} // 加-f 则 读到文件EOF时不操作，等待文件内容变化
			// fmt.Println("Lines: ", lineNum) // debug,定稿时必须去掉
			time.Sleep(time.Second * 1) // 1秒延迟可以阻止 lineNum 自增太快，很有必要加！1秒以上的延迟会使得等待时间久
		case nil:
			// err为空
			if lineNum-1 <= countLines-lastNLines {
				// fmt.Println("lineNum:", lineNum, ", lastNLines:", lastNLines) // debug
				//对于 lastNLines 之前的行，直接跳过循环
				continue
			}
			fmt.Printf(line)
			// lastNLines--
		default:
			fmt.Println("错误：", err.Error())
			break
		}
	}
}

func main() {
	var (
		fpath        string
		follow, help bool
		lastNLines   int
	)

	flag.StringVar(&fpath, "t", "", "target file path")
	flag.BoolVar(&follow, "f", false, "output appended data as the file grows, press Ctrl+c to stop")
	flag.BoolVar(&help, "h", false, "Help Info")
	flag.IntVar(&lastNLines, "n", 10, "output the last x lines")

	flag.Usage = func() {
		fmt.Println("Usage: tail [OPTION]... [FILE]...")
		flag.PrintDefaults()
	}

	flag.Parse()

	if fpath == "" || help {
		flag.Usage()
		return
	}
	tail(lastNLines, fpath, follow)
}
