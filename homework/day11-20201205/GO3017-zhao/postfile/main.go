package main

import (
	"fmt"
	"net/http"
	"zhao/db"
	_ "zhao/routes"
)

func main() {
	driverName := "mysql"
	// dsn := "root:zhaO..123@tcp(172.16.212.137:3306)/postfile?charset=utf8mb4&loc=Local&parseTime=true"
	dsn := "test:NDg3NTBi@tcp(127.0.0.1:3306)/postfile?charset=utf8mb4&loc=Local&parseTime=true"
	err := db.OpenDB(driverName, dsn)
	if err != nil {
		fmt.Println("1111", err)
		return
	}
	defer db.CloseDB()
	fmt.Println(db.DB)

	err = http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
	}
}
