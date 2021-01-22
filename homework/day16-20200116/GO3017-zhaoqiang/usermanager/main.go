package main

import (
	"fmt"
	"net/http"
	"runtime"
	"zhao/config"
	"zhao/monitoring"
	_ "zhao/routers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	addr := ":8888"
	driverName := "mysql"
	dsn := "root:zhaO..123@tcp(zhao:3306)/usermanager?parseTime=true&loc=Local&charset=utf8mb4"
	err := config.OpenDb(driverName, dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer config.CloseDB()

	go func() {
		monitoring.GorouterNumber.Set(float64(runtime.NumGoroutine()))
	}()

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
