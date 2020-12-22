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
	fmt.Println("main")
	var driverName = "mysql"
	var dsn = "web:web@tcp(127.0.0.1:3306)/user?parseTime=true&loc=Local&charset=utf8mb4"
	if err := models.InitDB(driverName, dsn); err != nil {
		fmt.Println(err)
		return
	}
	defer models.CloseDB()
	routes.Route()
	fmt.Println("server started at: ", addr)
	http.ListenAndServe(addr, nil)

	//initSql := `
	//INSERT INTO user
	//(name, password, sex, born, address, cell, created_at, updated_at)
	//  VALUES
	//('admin', password('admin123'), 1, '1995.03.04', 'Venus', '18811739999', NOW(), NOW()),
	//('jaccy', password('jaccy'), 1, '1895.06.04', 'London', '18811738998', NOW(), NOW()),
	//('leslie', password('imwhatim'), 1, '1975.03.04', 'HongKong', '18811738888', NOW(), NOW());
	//`

	//models.DB.Exec(initSql)
	//err := services.ListAllUser(models.DB)
	//errf := services.IDFindUser(models.DB, 1)
	//errn := services.NameFindUser(models.DB, "leslie")
	//if err != nil || errf != nil || errn != nil {
	//	fmt.Println(err, errf, errn)
	//	return
	//}
	//services.IDDelUser(models.DB, 3)
	//fmt.Println("===")
	//services.ListAllUser(models.DB)
	//fmt.Println("===")
	//if err := services.NameDelUser(models.DB, "jack"); err != nil {
	//	fmt.Printf("deleted user: %v", "jack")
	//}
	//services.ListAllUser(models.DB)
	fmt.Println("------------")
	id, err := services.GetMaxID(models.DB)
	fmt.Println(id, err)
	fmt.Println("main done")

}
