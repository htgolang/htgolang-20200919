package models

import (
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cast"
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

func DbToCsv() {
	newFile, err := os.Create("res/test.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		newFile.Close()
	}()
	// 写入UTF-8
	newFile.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM，防止中文乱码
	// 写数据到csv文件
	w := csv.NewWriter(newFile)
	header := []string{"ID", "Name", "Tel", "Addr", "Birthday", "Passwd"} //标题
	data := [][]string{
		header,
	}
	for _, v := range UsersDb {
		context := []string{
			cast.ToString(v.ID),
			v.Name,
			v.Tel,
			v.Addr,
			cast.ToString(v.Birthday),
			v.Passwd,
		}

		data = append(data, context)

	}
	w.WriteAll(data)
	w.Flush()
}

func CsvToDb() {
	f, err := os.Open("res/test.csv")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	reader := csv.NewReader(f)

	// 可以一次性读完
	result, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	// fmt.Println(result)
	for _, v := range result {
		user := &Users{
			ID:       cast.ToInt(v[0]),
			Name:     v[1],
			Addr:     v[2],
			Tel:      v[3],
			Birthday: cast.ToTime(v[4]),
			Passwd:   v[5],
		}
		UsersDb[v[1]] = user
	}
	delete(UsersDb, "Name")
	fmt.Println(UsersDb)
}

func DbToGob() {
	file, err := os.Create("res/test.gob")
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range UsersDb {
		UserDb = append(UserDb, v)
		// fmt.Println(UserDb)
	}
	enc := gob.NewEncoder(file)
	if err := enc.Encode(UserDb); err != nil {
		fmt.Println(err)
	}

}

func GobToDb() {
	File, _ := os.Open("res/test.gob")
	D := gob.NewDecoder(File)
	D.Decode(&UserDb)
	for _, v := range UserDb {
		UsersDb[v.Name] = v
	}
	// fmt.Println(UsersDb)
}
