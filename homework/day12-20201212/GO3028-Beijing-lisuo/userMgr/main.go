package main

import (
	"fmt"
	"userMgr/models"
	_ "userMgr/routers"

	beego "github.com/astaxie/beego"
)

func main() {
	var driverName = "mysql"
	// var dsn = "web:web@tcp(127.0.0.1:3306)/user?parseTime=true&loc=Local&charset=utf8mb4"
	dsn := "test:NDg3NTBi@tcp(127.0.0.1:3306)/testboy?parseTime=true&loc=Local&charset=utf8mb4"	
	if err := models.InitDB(driverName, dsn); err != nil {
		fmt.Println(err)
		return
	}
	defer models.CloseDB()
	beego.Run()
}
