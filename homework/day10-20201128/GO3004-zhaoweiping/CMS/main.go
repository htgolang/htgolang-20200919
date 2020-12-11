package main

import (
	"CMS/controllers"
	"CMS/models"
	"CMS/routers"
	"CMS/utils"
	"net/http"
)

func ChoiceInit() {
	// choice := utils.ChoiceYesOrNo("是否要重新初始化并重置数据，添加用户?(Y/n)：")
	choice := true
	if choice {
		// for i := 0; i < 3; i++ {
		// 	if controllers.AddUserInit() {
		// 		return
		// 	}
		// }
		return
	}
}
func InitDbType(dbType string) {
	controllers.DbType = dbType
	// fmt.Println("你可以选择初始化工具 init，使用：go run main.go -h 查看具体细节，默认使用 csv 存储数据，可以使用 go run main -i -t csv/gob/json 手动替换初始化")

	switch dbType {
	case "csv":
		// fmt.Println("csv")
		models.CsvToDb()
		ChoiceInit()
	case "gob":
		// fmt.Println("gob")
		models.GobToDb()
		ChoiceInit()
	case "json":
		// fmt.Println("json")
		models.JsonToDb()
		ChoiceInit()
	}
}

func main() {
	dbType := utils.FlagMsg()

	InitDbType(dbType)

	// utils.PrintMsg()
	addr := ":8080"
	routers.Register()
	// manager.Run()
	http.ListenAndServe(addr, nil)
}
