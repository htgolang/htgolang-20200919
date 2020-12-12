package config

import (
	"database/sql"
)

var Db *sql.DB

func InitDb(driverName, dsn string) error {
	var err error
	Db, err = sql.Open(driverName, dsn)

	if err != nil {
		return err
	}
	if err := Db.Ping(); err != nil {
		return err
	}
	return nil
}

func CloseDb() error {
	return Db.Close()
}
