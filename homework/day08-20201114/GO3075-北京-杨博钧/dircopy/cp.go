package main

import (
	"dircopy/model"
	"dircopy/utils"
	"fmt"
	"sync"
)

func main() {
	// wg用来等待工作协程完成
	var wg sync.WaitGroup

	// 初始化文件夹拷贝
	file := model.NewCopyFile()

	// 检查传入路径是否符合条件
	// 当符合条件时触发拷贝函数
	isdocopy, err := file.Check()
	if err != nil {
		fmt.Println(err)
	}
	if isdocopy {
		utils.Copy(file, &wg)
	}

	// 等待工作协程完成
	wg.Wait()
}