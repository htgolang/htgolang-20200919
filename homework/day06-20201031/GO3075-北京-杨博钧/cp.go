package main

import (
	"fmt"
	"go_project/homework/day6/GO3075-北京-杨博钧/utils"
)
//作业一:拷贝文件
func main() {
	copyFile := utils.NewCopyFile()
	err := copyFile.Copy()
	if err != nil {
		fmt.Printf("拷贝失败:%v\n", err)
	}
}
