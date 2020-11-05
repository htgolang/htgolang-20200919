package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// 两个文件指针 从sf读 覆盖写入df
func overWrite(sf, df *os.File) {
	err := df.Truncate(0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		content := make([]byte, 20)
		for {
			n, err := sf.Read(content)
			if err != io.EOF {
				df.Write(content[:n])
			} else {
				break
			}
		}
	}
}

func overWrite2(sf, df *os.File) {
	err := df.Truncate(0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		content := make([]byte, 20)
		for {
			n, err := sf.Read(content)
			if err != io.EOF {
				df.Write(content[:n])
			} else {
				break
			}
		}
	}
}

// src路径 dst路径
func copyFile(srcFile, dstFile string) {
	// fmt.Println(srcFile, dstFile)
	sf, err := os.OpenFile(srcFile, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件错误: %v \n", err)
		return
	}
	defer sf.Close()
	//检查dst文件是否存在
	_, err2 := os.Stat(dstFile)
	// 如果dst不存在就创建
	if err2 != nil {
		df, _ := os.OpenFile(dstFile, os.O_CREATE, 0644)
		defer df.Close()
		overWrite(sf, df)
		fmt.Println("CP执行成功")
	} else {
		df, _ := os.OpenFile(dstFile, os.O_TRUNC|os.O_WRONLY, 0644)
		defer df.Close()
		//如果存在就提示覆盖
		fmt.Print(dstFile, "已经存在,是否覆盖? y/N : ")
		var overwrite string
		fmt.Scanf("%s", &overwrite)
		switch overwrite {
		case "y", "Y":
			overWrite(sf, df)
			fmt.Println("文件覆盖成功")
		default:
			fmt.Println("不执行覆盖,cp失败")
		}
		return
	}
}
func main() {
	var (
		srcFile, dstFile string
		help             bool
	)
	flag.StringVar(&srcFile, "s", "", "src file path")
	flag.StringVar(&dstFile, "d", "", "dst file path")
	flag.BoolVar(&help, "h", false, "Help Info")
	flag.Usage = func() {
		fmt.Println(`
Usage: cp.exe -s src -d dst
Options:`)
		flag.PrintDefaults()
	}

	flag.Parse()

	if help {
		flag.Usage()
	} else {
		copyFile(srcFile, dstFile)
	}
}
