package utils

import (
	"fmt"
	"os"
)

//判断一个用户是否存在,存在时返回切片下标
func isUserExists(name string) (int, bool) {
	for i, user := range UsersList {
		if name == user.Name {
			return i, true
		}
	}
	return 0, false
}

func HasDataDIr() {
	//默认保存在当前路径的data文件夹下，先检查是否存在，不存在则先创建
	stat, err := os.Stat("data")
	if os.IsNotExist(err) {
		os.Mkdir("data", 0666)
	} else if !stat.IsDir() {
		fmt.Println("存在名为data文件，无法创建保存的目录")
		return
	}
}