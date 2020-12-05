package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	types := map[string]string{
		"BIGINT":  "int64",
		"VARCHAR": "string",
		"TEXT":    "string",
		"DATE":    "*time.Time",
		"TINYINT": "bool",
	}
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

	sql := `
		show tables
	`
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			fmt.Println(err)
			break
		} else {
			fieldSQL := fmt.Sprintf("select * from %s", tableName)
			fieldRows, err := db.Query(fieldSQL)
			if err != nil {
				fmt.Println(tableName, err)
			} else {
				columns, err := fieldRows.ColumnTypes()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("**************************")
					fmt.Println(tableName)
					file, _ := os.Create(fmt.Sprintf("%s.go", tableName))
					fmt.Fprintf(file, "type %s struct {\n", tableName)

					for _, column := range columns {
						fmt.Println("\t", column.Name(), column.DatabaseTypeName())
						fmt.Fprintf(file, "\t%s %s\n", column.Name(), types[column.DatabaseTypeName()])
					}

					fmt.Fprintf(file, "}\n")

					// type name struct {

					//}
				}
			}
			fieldRows.Close()
		}
	}
}
