package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	driverName := "mysql"
	dsn := "golang:golang@2020@tcp(10.0.0.2:3306)/user?parseTime=true&loc=Local&charset=utf8mb4"
	db, err := sql.Open(driverName, dsn)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	name, password, birthday := "kk2", "xxxx", "1999-11-11"
	sql := `
	INSERT INTO user(name, password, birthday) VALUES(?, ?, ?);
	`
	result, err := db.Exec(sql, name, password, birthday)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	}

	sql = `
	UPDATE user
	SET birthday=now()
	WHERE id=?
	`
	result, err = db.Exec(sql, 1)
	fmt.Println(err)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

	sql = `
	DELETE from user where name=?
	`

	result, err = db.Exec(sql, name)
	fmt.Println(err)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
