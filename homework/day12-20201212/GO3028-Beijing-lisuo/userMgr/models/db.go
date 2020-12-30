package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB        *sql.DB
	errOpenDB error
)

// InitDB open a db connection pool
func InitDB(driverName, dsn string) error {
	DB, errOpenDB = sql.Open(driverName, dsn)
	if errOpenDB != nil {
		return errOpenDB
	}
	if err := DB.Ping(); err != nil {
		return err
	}
	fmt.Printf("opened db: %#v\n", (*DB))
	return nil
}

// CloseDB returns the connection to the connection pool, rarely necessary
func CloseDB() error {
	return (*DB).Close()
}
