package main

import (
	"fmt"
	"net/http"

	"user/config"
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
	http.ListenAndServe(addr, nil)
}
