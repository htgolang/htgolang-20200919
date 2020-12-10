package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func OpenDB(drivername, dsn string) error {
	var err error
	DB, err = sql.Open(drivername, dsn)
	if err != nil {
		fmt.Println(err, "sql Open database connect false")
		return err
	}
	if err := DB.Ping(); err != nil {
		fmt.Println(err, "db ping false")
		return err
	}
	return nil
}

func CloseDB() error {
	return DB.Close()
}
