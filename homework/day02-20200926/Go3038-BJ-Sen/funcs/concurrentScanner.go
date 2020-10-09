package funcs

import (
	"bufio"
	"fmt"
	"os"
)

const sBufferSize = 100

type sChunk struct {
	bufsize int
	offset int64
}

func ConcurrentScanner() {
	file, err := os.Open("day02/I have a dream")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for {
		read := scanner.Scan()
		if !read {
			break
		}
		fmt.Println("read byte array:", scanner.Bytes())
		fmt.Println("read string:", scanner.Text())
	}
}
