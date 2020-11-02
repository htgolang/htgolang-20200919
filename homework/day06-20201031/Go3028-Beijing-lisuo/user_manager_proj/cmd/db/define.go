package db

import "github.com/htgolang/htgolang-20200919/tree/master/homework/day06-20201031/Go3028-Beijing-lisuo/user_manager_proj/define"

// dbLocation --> basedir/db/userDB.csv
var (
	dbLocation string
	absDir     string
	dbAbsFile  string
	dbDir      = "db"
	dbName     = "userDB.csv"
	base       = "main.go"
)

// contains users read from dbNam
var tmpUsers []define.User
