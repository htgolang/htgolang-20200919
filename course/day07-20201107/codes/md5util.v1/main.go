package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

const bufferSize = 1024

func main() {
	var path string

	flag.StringVar(&path, "path", "", path)

	flag.Usage = func() {
		fmt.Println("usage: md5util --path path")
	}

	flag.Parse()

	// 检查

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	hasher := md5.New()

	// io.Copy(dst, src)

	ctx := make([]byte, bufferSize)
	for {
		n, err := file.Read(ctx)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		hasher.Write(ctx[:n])
	}

	fmt.Printf("%x\n", hasher.Sum(nil))
}
