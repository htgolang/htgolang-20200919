package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const bufferSize = 1024

func main() {
	var (
		src string
		dst string
	)

	flag.StringVar(&src, "src", "", "src file path")
	flag.StringVar(&dst, "dst", "", "dst file path")

	flag.Usage = func() {
		fmt.Println("usage: cp --src path --dst path")
	}

	flag.Parse()

	// 检查
	/* src
	没输入
	文件路径不存在
	目录

	dst
	没输入
	父目录不存在 => 自动创建父目录（创建失败/创建成功），报错
	文件存在
	目录
	*/

	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dstFile.Close()

	ctx := make([]byte, bufferSize)

	for {
		n, err := srcFile.Read(ctx)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		dstFile.Write(ctx[:n])
	}
}
