package utils

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"
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
	timeNow := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("data/%v.gob", timeNow)
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
	DelFile()
}
//定义读取函数
func GobLoad(name string) {
	//先清空当前用户信息
	UsersList = []model.User{}
	file, err := os.Open(name)
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