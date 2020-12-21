package main

import (
	"fmt"
	"user_manager/models"
	"user_manager/services"
)

func main() {
	fmt.Println("main")
	var driverName = "mysql"
	var dsn = "web:web@tcp(127.0.0.1:3306)/user?parseTime=true&loc=Local&charset=utf8mb4"
	if err := models.InitDB(driverName, dsn); err != nil {
		fmt.Println(err)
		return
	}
	defer models.CloseDB()

	err := services.ListAllUser(models.DB)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("main done")
}
