package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func flagMsg() string {
	//定义命令行参数方式1
	var srcFile string
	flag.StringVar(&srcFile, "s", "", "源文件")
	//解析命令行参数
	flag.Parse()

	if srcFile == "" {
		flag.Usage()
		os.Exit(1)
	}
	return srcFile
}
func listDir(path string, indent int) {
	//检测地址是否是绝对地址
	dir, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("err : %s \n", err)
		return
	}
	// fmt.Printf("%sDir: %s\n", strings.Repeat(" ", indent*4), dir)
	finfos, err := ioutil.ReadDir(dir)
	if err != nil {
		f, err := os.Stat(path)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(f.Name())
		return
	}
	// 遍历子目录或者文件
	for _, fi := range finfos {
		// 如果是目录，则递归输出
		if fi.IsDir() {
			listDir(dir+string(os.PathSeparator)+fi.Name(), indent+1)
			continue
		}
		// 如果是文件，则直接输出文件名
		fmt.Printf("%s%s\n", strings.Repeat(" ", indent*4), fi.Name())
	}
}
func main() {
	srcFile := flagMsg()
	listDir(srcFile, 0)
}
