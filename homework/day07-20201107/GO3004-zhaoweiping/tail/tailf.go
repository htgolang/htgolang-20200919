package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

//命令行参数解析
func FlagMsg() string {
	//定义命令行参数方式1
	var (
		follow string
	)
	flag.StringVar(&follow, "f", "", "filepath")

	//解析命令行参数
	flag.Parse()
	return follow
}

func main() {
	filepath := FlagMsg()
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Open file fail:%v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// time.Sleep(1 * time.Second)
				continue
			} else {
				break
			}
		}
		fmt.Println(string(line))
	}
}
