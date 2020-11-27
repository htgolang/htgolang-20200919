package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

//命令行参数解析
func flagMsg() (string, string) {
	//定义命令行参数方式1
	var (
		srcFile string
		dstFile string
	)
	flag.StringVar(&srcFile, "s", "", "源文件")
	flag.StringVar(&dstFile, "d", "", "目的文件")
	//解析命令行参数
	flag.Parse()

	if srcFile == "" && dstFile == "" {
		flag.Usage()
		os.Exit(1)
	}
	return srcFile, dstFile
}

//读取源文件
func srcf(srcFile string) string {
	sFile, err := os.Open(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	defer sFile.Close()
	ctx := make([]byte, 1024)
	for {
		n, err := sFile.Read(ctx)
		if err != nil {
			log.Fatal(err)
		}
		return string(ctx[:n])
	}
}

// 判断所给路径文件/文件夹是否存在，存在则返回 true
func existsFile(dstFile string) bool {
	_, err := os.Stat(dstFile) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 若目标文件存在，是否选择强制覆盖，是就返回 true
func forceCpFile(dstFile string) bool {
	existsFileBool := existsFile(dstFile)
	if existsFileBool == true {
		confirm := Input("请输入 y|n 确认是否强制覆盖已存在的文件：")
		if strings.ToLower(confirm) == "y" || strings.ToLower(confirm) == "yes" {
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}

// 扫描输入
func Input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return text
}

func writeFile(srcFileByte, dstFile string) {
	dFile, err := os.Create(dstFile)
	if err != nil {
		log.Fatal(err)
	}
	defer dFile.Close()
	dFile.WriteString(srcFileByte)
}

//拷贝文件函数
func copyFile(srcFileByte, dstFile string) {
	dstFileBool := forceCpFile(dstFile)

	if dstFileBool == true {
		writeFile(srcFileByte, dstFile)
		fmt.Println("copy done!")
	}
}

func main() {
	srcFile, dstFile := flagMsg()
	srcFileByte := srcf(srcFile)
	copyFile(srcFileByte, dstFile)
}
