package models

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cast"
)

var path = "./res/csv/"
var name = "user.csv"
var FilePath = DealWithFiles(path, name)

func DbToCsv() {

	newFile, err := os.Create(FilePath)
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
	header := []string{"ID", "Name", "Sex", "Addr"} //标题
	w.Write(header)
	for _, v := range UsersDb {
		context := []string{
			cast.ToString(v.ID),
			v.Name,
			cast.ToString(v.Sex),
			v.Addr,
		}
		// data = append(data, context)
		w.Write(context)
	}

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
	w.Flush()
	// data := [][]string{
	// 	header,
	// }
	// fmt.Println(UsersDb)
	// for _, v := range UsersDb {
	// 	context := []string{
	// 		cast.ToString(v.ID),
	// 		v.Name,
	// 		v.Tel,
	// 		v.Addr,
	// 		cast.ToString(v.Birthday),
	// 		v.Passwd,
	// 	}
	// 	// fmt.Println(context)
	// 	data = append(data, context)
	// 	// fmt.Println(data)

	// }
	// w.WriteAll(data)
	// w.Flush()
}

func CsvToDb() {

	f, err := os.Open(FilePath)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer f.Close()
	reader := csv.NewReader(f)

	// 可以一次性读完
	result, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	// fmt.Println(result[0][0], "ok")
	// if result[0][0] != " " {
	// 	fmt.Println(result)
	// }
	for _, v := range result {
		user := &Users{
			ID:   cast.ToInt(v[0]),
			Name: v[1],
			Sex:  cast.ToBool(v[2]),
			Addr: v[3],
		}
		UsersDb[v[1]] = user
	}

	delete(UsersDb, "Name")
	// fmt.Println(UsersDb)
}
