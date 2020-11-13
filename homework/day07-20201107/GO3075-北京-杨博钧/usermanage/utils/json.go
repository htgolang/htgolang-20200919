package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"encoding/json"
	"usermanage/model"
)

type JsonPersist struct {}

func (this JsonPersist) Save() {
	timeNow := time.Now().Format("20060102_150405")
	fileName := fmt.Sprintf("data/%v.json", timeNow)
	file, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer file.Close()
	ctx, err := json.Marshal(UsersList)
	if err != nil {
		fmt.Printf("json编码失败:%v", err)
		return
	}
	file.Write(ctx)
	DelFile()
}

func JsonLoad(name string) {
	ctx, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Printf("文件读取失败:%v", err)
		return
	}
	UsersList = []model.User{}
	err = json.Unmarshal(ctx, &UsersList)
	if err != nil {
		fmt.Printf("json解析失败:%v", err)
		return
	}
}