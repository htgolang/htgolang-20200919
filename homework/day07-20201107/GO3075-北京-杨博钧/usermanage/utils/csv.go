package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
	"usermanage/model"
)

type CsvPersister struct {
}

//定义保存函数
func (this CsvPersister) Save() {
	//以只写模式打开文件，包含清空、创建属性
	timeNow := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("data/%v.csv", timeNow)
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//开启编码工具
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, v := range UsersList {
		writer.Write([]string{strconv.Itoa(v.Id), v.Name, v.Addr, v.Tel,
			v.Birthday.Format("2006-01-02"), v.Password})
	}
	DelFile()
}
//定义读取函数
func CsvLoad(name string) {
	//先清空当前用户信息
	UsersList = []model.User{}
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("没有可读取数据文件")
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		string, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		id, _ := strconv.Atoi(string[0])
		birthday, _ := time.Parse("2006-01-02", string[4])
		UsersList = append(UsersList, model.User{id, string[1], string[2],
			string[3], birthday, string[5]})
	}
}