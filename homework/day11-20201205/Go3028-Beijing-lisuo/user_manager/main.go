package main

import (
	"fmt"
	"net/http"
	"user_manager/models"
	"user_manager/routes"
	"user_manager/services"
)

var (
	addr        = ":8889"
	templateAbs = "/data/htgolang-20200919/homework/day11-20201205/Go3028-Beijing-lisuo/user_manager"
)

func main() {
	var driverName = "mysql"
	// var dsn = "web:web@tcp(127.0.0.1:3306)/user?parseTime=true&loc=Local&charset=utf8mb4"
	dsn := "test:NDg3NTBi@tcp(127.0.0.1:3306)/usermanager?parseTime=true&loc=Local&charset=utf8mb4"	
	if err := models.InitDB(driverName, dsn); err != nil {
		fmt.Println(err)
		return
	}
	defer models.CloseDB()
	routes.Route()
	fmt.Println("server started at: ", addr)
	http.ListenAndServe(addr, nil)

	fmt.Println("------------")
	id, err := services.GetMaxID(models.DB)
	fmt.Println(id, err)
	fmt.Println("main done")

}
