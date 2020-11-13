package utils

import (
	"encoding/gob"
	"fmt"
	"os"
	"usermanage/model"
)

type GobPersister struct {
}
//将结构体注册到gob中
func init() {
	gob.Register(model.User{})
}
//定义保存函数
func (this GobPersister) Save() {
	//默认保存在当前路径的data文件夹下，先检查是否存在，不存在则先创建
	stat, err := os.Stat("data")
	if os.IsNotExist(err) {
		os.Mkdir("data", 0666)
	} else if !stat.IsDir() {
		fmt.Println("存在名为data文件，无法创建保存的目录")
		return
	}
	//以只写模式打开文件，包含清空、创建属性
	filename := "data/data.gob"
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//开启编码工具
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(UsersList)
	if err != nil {
		fmt.Println(err)
		return
	}
}
//定义读取函数
func (this GobPersister) Load() {
	//先清空当前用户信息
	UsersList = []model.User{}
	filename := "data/data.gob"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("没有可读取数据文件")
		return
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&UsersList)
	if err != nil {
		fmt.Println(err)
		return
	}
}