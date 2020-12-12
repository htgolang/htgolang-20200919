package main

import (
	"fmt"

	"user/config"
	_ "user/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	driverName := "mysql"
	dsn := "golang:golang@2020@tcp(10.0.0.2:3306)/user?parseTime=true&loc=Local&charset=utf8mb4"

	if err := config.InitDb(driverName, dsn); err != nil {
		fmt.Println(err)
		return
	}
	defer config.CloseDb()

	addr := ":8888"

	beego.Run(addr)
}
