package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"time"
	"usermanage/model"
)

// 创建文件列表类型用于按照时间排序
type MyFileInfoList []os.FileInfo

func (this MyFileInfoList) Len() int {
	return len(this)
}

func (this MyFileInfoList) Less(i, j int) bool {
	return this[j].ModTime().Sub(this[i].ModTime()) > 0
}

func (this MyFileInfoList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

// 持久化数据函数
func MkDataDIr() {
	// 默认保存在当前路径的data文件夹下，先检查是否存在，不存在则先创建
	stat, err := os.Stat("data")
	if os.IsNotExist(err) {
		os.Mkdir("data", 0666)
	} else if !stat.IsDir() {
		fmt.Println("存在名为data文件，无法创建保存的目录")
		return
	}
}

// 文件按照修改时间排序
func GetSortedFileList() MyFileInfoList {
	var fileInfos MyFileInfoList
	path, _ := os.Open("data")
	files, _ := path.Readdir(-1)
	if len(files) == 0 {
		return []os.FileInfo{}
	}
	for _, file := range files {
		fileInfos = append(fileInfos, file)
	}
	sort.Sort(fileInfos)
	return fileInfos
}

// 定义只保留最后5次记录函数
func DelFile() {
	fileInfos := GetSortedFileList()
	if len(fileInfos) > 5 {
		for _, file := range fileInfos[:len(fileInfos) - 5] {
			os.Remove("data/" + file.Name())
		}
	}
}

// 定义保存数据和加载数据函数
func SaveData() {
	// 先检查是否有data目录没有则创建
	MkDataDIr()

	// 定义文件保存时的名称
	timeNow := time.Now().Format("20060102_150405")
	fileName := fmt.Sprintf("data/%v.json", timeNow)

	// 写入数据到文件
	file, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer file.Close()
	ctx, err := json.Marshal(UsersList)
	if err != nil {
		fmt.Printf("json编码失败:%v", err)
		return
	}
	file.Write(ctx)

	// 删除超过5次的保存记录
	DelFile()
}

func LoadData() {
	// 创建保存路径
	MkDataDIr()

	// 获取路径下所有文佳
	fileInfos := GetSortedFileList()

	// 判断当前为空时则直接返回
	if len(fileInfos) == 0 {
		return
	}

	// 加载最后一个文件的数据
	file := fileInfos[len(fileInfos) - 1]
	fileName := fmt.Sprintf("data/%v", file.Name())
	ctx, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("文件读取失败:%v", err)
		return
	}
	UsersList = []*model.User{}
	err = json.Unmarshal(ctx, &UsersList)
	if err != nil {
		fmt.Printf("json解析失败:%v", err)
		return
	}
}