package models

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB        *sql.DB
	errOpenDB error
)

// init get db ready
func init() {
	dsn := beego.AppConfig.String("db::dsn")
	orm.RegisterModel(new(User))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	if err := orm.RegisterDataBase("default", "mysql", dsn); err != nil {
		panic(err)
	}
	orm.RunSyncdb("default", false, true)
}

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
