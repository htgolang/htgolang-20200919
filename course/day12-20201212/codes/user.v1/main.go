package main

import (
	"fmt"
	"net/http"
	"os"

	"user/config"
	"user/logger"
	"user/routers"

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

	addr := ":8080"

	routers.Register()

	logFile, _ := os.OpenFile("web.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)

	logger.InitLogger(logFile)
	http.ListenAndServe(addr, nil)
}
