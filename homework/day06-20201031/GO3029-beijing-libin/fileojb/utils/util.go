package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//对目标文件进行写操作的函数
func HandleWriteFile(src *os.File, filename string, flag int, mode os.FileMode) error {
	fileobj, err := os.OpenFile(filename, flag, mode)
	if err != nil {
		return err
	}
	defer fileobj.Close()
	_, err = io.Copy(fileobj, src)
	if err != nil {
		fmt.Println("copy过程失败，原因是", err.Error())
		return err
	}
	return nil
}

//检测文件是否存在
func CheckFileExist(filename string) bool {
	//首先，检测文件或者目录是否存在
	fileobj, err := os.Stat(filename)
	if err != nil {
		return false
	}

	// 判断filename若是目录
	if fileobj.IsDir() {
		//暂时不处理目录
		fmt.Println("is directory")
		return false
	}
	return true
}

//检测目标为目录且存在
func CheckDirExist(dirname string) bool {
	dirInfo, err := os.Stat(dirname)
	if err != nil {
		fmt.Printf("%s不存在\n")
		return false
	}
	if dirInfo.IsDir() {
		return true
	}
	return false
}

//返回文件的权限
func GetFilePerm(filename string) os.FileMode {
	fileInfo, _ := os.Stat(filename)
	return fileInfo.Mode().Perm()
}

//返回一个文件的行数
func GetFileLines(filename string) (int, error) {
	var count int
	fileobj, err := os.Open(filename)
	if err != nil {
		fmt.Println("打开文件失败，原因是", err.Error())
	}
	defer fileobj.Close()
	filereader := bufio.NewReader(fileobj)
	for {
		_, err := filereader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("读取文件行出现错误，原因是", err.Error())
			}
			return count, err
		}
		count++
	}
	return count, nil

}
