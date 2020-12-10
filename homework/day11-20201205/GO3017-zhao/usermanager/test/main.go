package main

import (
	"database/sql"
	"fmt"
	"time"
	"zhao/models"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	driverName := "mysql"
	dsn := "root:zhaO..123@tcp(172.16.212.137:3306)/usermanager?loc=Local&parseTime=true&charset=utf8"
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select id, name, sex, addr, tel, brithday, passwd, create_at from user")
	if err != nil {
		fmt.Println("[server.GetUsers.Query]", err) //如果数据库出错返回一个空用户列表
		return
	}
	defer rows.Close()
	users := make([]*models.User, 0, 20)
	for rows.Next() {
		var (
			id         int64
			name       string
			sex        bool
			addr       string
			tel        string
			brithday   *time.Time
			passwd     string
			createTime *time.Time
		)
		err := rows.Scan(&id, &name, &sex, &addr, &tel, &brithday, &passwd, &createTime)
		if err != nil {
			fmt.Println("[server.GetUSers.Scan]", err)
			return
		}
		users = append(users, &models.User{id, name, sex, addr, tel, brithday, passwd, createTime})
	}
	for _, user := range users {
		fmt.Println(user)
	}
}
