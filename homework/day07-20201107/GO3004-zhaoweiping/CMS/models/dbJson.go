package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func DbToJson() {
	// fmt.Println(UsersDb)
	b, err := json.Marshal(UsersDb)
	if err != nil {
		fmt.Println(err)
	}
	err1 := ioutil.WriteFile("res/test.json", b, os.ModeAppend)
	if err1 != nil {
		fmt.Println(err1)
	}
}
func JsonToDb() {
	filePtr, err := os.Open("res/test.json")
	if err != nil {
		fmt.Println("文件打开失败 [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()
	decoder := json.NewDecoder(filePtr)
	decoder.Decode(&UsersDb)
	// fmt.Println(UsersDb)
}
