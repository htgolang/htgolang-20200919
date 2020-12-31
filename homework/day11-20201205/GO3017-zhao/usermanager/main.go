package main

import (
	"fmt"
	"net/http"
	"zhao/config"
	_ "zhao/routers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	addr := ":8888"
	driverName := "mysql"
	// dsn := "root:zhaO..123@tcp(172.16.212.137:3306)/usermanager?parseTime=true&loc=Local&charset=utf8mb4"
	dsn := "test:NDg3NTBi@tcp(127.0.0.1:3306)/usermanager?parseTime=true&loc=Local&charset=utf8mb4"	
	err := config.OpenDb(driverName, dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer config.CloseDB()

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
