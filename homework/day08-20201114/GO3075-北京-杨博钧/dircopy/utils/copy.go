package utils

import (
	"bufio"
	"dircopy/model"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)
// 定义文件拷贝函数
func CopyFile(src, dest string, wg *sync.WaitGroup) {
	// 关闭一个等待
	defer wg.Done()

	// 开启一个读文件的句柄
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)
	mode := os.O_CREATE

	// 判断目标文件是否存在，如果存在需要用户进行判断是否替换
	_, err = os.Stat(dest)
	if err == nil {
		var choise string
	loop:
		for {
			fmt.Printf("%v文件已经存在是否进行覆盖?(Y/N):\n", dest)
			fmt.Scan(&choise)
			switch choise {
			case "Y", "y" :
				mode = os.O_TRUNC
				break loop
			case "N", "n" :
				fmt.Println("文件已存在不进行替换\n")
				return
			default :
				fmt.Println("请输入正确的选项....")
			}
		}
	} else if !os.IsNotExist(err) {
		fmt.Println(err)
		return
	}

	// 打开写的文件句柄
	desFile, err := os.OpenFile(dest, mode | os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer desFile.Close()
	writer := bufio.NewWriter(desFile)
	defer writer.Flush()

	// 使用io.Copy完成文件拷贝
	fmt.Printf("正在拷贝%v到%v\n", src, dest)
	io.Copy(writer, reader)
}
// 拷贝文件夹函数
func Copy(info *model.CopyDir, wg *sync.WaitGroup) {
	// 判断目标文件夹是否存在,存在则直接创建,否则
	_, err := os.Stat(info.DestName)
	if os.IsNotExist(err) {
		os.Mkdir(info.DestName, 0666)
	} else {
		os.MkdirAll(info.DestName + "/" +
			strings.Split(info.SrcName, "/")[len(strings.Split(info.SrcName, "/")) -1],
			0666)
		info.DestName = info.DestName + "/" + strings.Split(info.SrcName, "/")[len(strings.Split(info.SrcName, "/")) -1]
	}

	// 遍历文件夹判断是文件还是目录
	file, _ := os.Open(info.SrcName)
	fileInfos, _ := file.Readdir(-1)
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			// 对于文件开启一个工作协程进行拷贝
			wg.Add(1)
			go CopyFile(info.SrcName + "/" + fileInfo.Name(),
				info.DestName + "/" + fileInfo.Name(), wg)
		} else {
			// 对于目录进行递归调用
			newInfo := &model.CopyDir{
				SrcName: info.SrcName + "/" + fileInfo.Name(),
				DestName: info.DestName,
			}
			Copy(newInfo, wg)
		}
	}
}