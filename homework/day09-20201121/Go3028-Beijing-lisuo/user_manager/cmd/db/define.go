package db

import (
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day09-20201121/Go3028-Beijing-lisuo/user_manager/define"
)

// dbLocation --> basedir/db/userDB.csv
var (
	dbLocation string
	absDir     string
	dbAbsFile  string
	base       = "main.go"
	dbDir      = "db"
	dbNameCSV  = "userDB.csv"
	dbNameGob  = "userDB.gob"
	dbNameJSON = "userDB.json"
	SaveFlag   = ""
)

var dbNameList map[string]string = map[string]string{
	"csv":  "userDB.csv",
	"gob":  "userDB.gob",
	"json": "userDB.json",
}

// contains users read from dbNam
var tmpUsers []define.User
