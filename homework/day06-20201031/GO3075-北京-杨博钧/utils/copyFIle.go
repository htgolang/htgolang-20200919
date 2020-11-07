package utils

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type CopyFile struct {
	SrcName string
	DestName string
}

func NewCopyFile() *CopyFile {
	var copyFile CopyFile
	flag.StringVar(&(copyFile.DestName), "d", "", "目标文件")
	flag.StringVar(&(copyFile.SrcName), "s", "", "源文件")
	flag.Parse()
	return &copyFile
}

func (this *CopyFile) String() string {
	return fmt.Sprintf("要拷贝的文件为:%v,目标文件为:%v\n", this.SrcName, this.DestName)
}

func (this *CopyFile) Copy() error {
	//判断源文件是否存在，不存在则报错返回
	_, err := os.Stat(this.SrcName)
	if os.IsNotExist(err) {
		return err
	} else if err != nil {
		return err
	}
	//打开文件句柄
	srcFile, err := os.Open(this.SrcName)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)
	mode := os.O_CREATE
	//判断目标文件是否存在，如果存在需要用户进行判断是否替换
	_, err = os.Stat(this.DestName)
	if err == nil {
		var choise string
	loop:
		for {
			fmt.Printf("%v文件已经存在是否进行覆盖?(Y/N):", this.DestName)
			fmt.Scan(&choise)
			switch choise {
			case "Y", "y" :
				mode = os.O_TRUNC
				break loop
			case "N", "n" :
				return fmt.Errorf("文件已存在不进行替换\n")
			default :
				fmt.Println("请输入正确的选项....")
			}
		}
	} else if !os.IsNotExist(err) {
		return err
	}
	//打开写的文件句柄
	desFile, err := os.OpenFile(this.DestName, mode | os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer desFile.Close()
	writer := bufio.NewWriter(desFile)
	defer writer.Flush()
	io.Copy(writer, reader)
	return nil
}
