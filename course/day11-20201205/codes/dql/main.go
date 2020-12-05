package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	driverName := "mysql"
	// "user:password@protocol(host:port)/dbname?charset=utf8mb4&loc=Local&parseTime=true"        //data store name 数据库连接信息，使用协议，用户&密码，数据库，连接参数

	dsn := "golang:golang@2020@tcp(10.0.0.2:3306)/user?charset=utf8mb4&loc=Local&parseTime=true" //data store name 数据库连接信息，使用协议，用户&密码，数据库，连接参数

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println(err)
		return
	}

	name := "%kk%" // sql注入
	sql := `
		select id, name, password, sex, birthday, addr, tel
		from user
		where name like ?
		order by ? desc
		limit ? offset ?
	`
	fmt.Println(sql)
	// 操作
	rows, err := db.Query(sql, name, "birthday", 3, 3) // 数据库的预处理方式
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id       int64
			name     string
			password string
			sex      bool
			birthday *time.Time
			addr     string
			tel      string
		)
		err := rows.Scan(&id, &name, &password, &sex, &birthday, &addr, &tel)
		if err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(id, name, password, sex, birthday, addr, tel)
		}
	}

	var id int64
	err = db.QueryRow("select id from user").Scan(&id)

	fmt.Println(err, id)
}
