package model

import (
	"flag"
	"fmt"
	"os"
)

type CopyDir struct {
	SrcName string
	DestName string
}

func NewCopyFile() *CopyDir {
	var copyFile CopyDir
	flag.StringVar(&(copyFile.DestName), "d", "", "目标文件夹")
	flag.StringVar(&(copyFile.SrcName), "s", "", "源文件夹")
	flag.Parse()
	return &copyFile
}
//标准化输出
func (this *CopyDir) String() string {
	return fmt.Sprintf("要拷贝的文件夹为:%v,目标文件夹为:%v\n", this.SrcName, this.DestName)
}
//检查源路径与目标路径是否符合要求
func (this *CopyDir) Check() (bool, error) {
	fmt.Println(this)
	stat, err := os.Stat(this.SrcName)
	if err != nil {
		return false, err
	}
	if !stat.IsDir() {
		return false, fmt.Errorf("拷贝源不是文件夹")
	}
	stat, err = os.Stat(this.DestName)
	if err != nil && !os.IsNotExist(err) {
		return false, err
	}
	if !os.IsNotExist(err) {
		if !stat.IsDir() {
			return false, fmt.Errorf("目标路径已存在并且并非文件夹")
		}
	}
	return true, nil
}