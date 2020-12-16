package process

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

var Db *sql.DB

type Fileinfo struct {
	Id int64
	Filename string
}
type TotalData struct {
	Datapie []map[string]string
	Databar []map[string]string
}

func Filemeta(filename string,size int64) (int64,error) {
	isql := "insert into fileinfo(filename, size, upload_at) values(?, ?, now())"
	r,err := Db.Exec(isql,filename,size)
	if err != nil {
		fmt.Println(err)
		return 0,err
	}
	id,err := r.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return 0,err
	}
	return id,nil

}

func Loginsert(id int64,line string) error  {
	sts := strings.Split(line," ")
	isql := "insert into recorder(fileid, ipaddr,method,status, insert_at) values(?, ?, ?, ?, now())"

	_,err := Db.Exec(isql,id,strings.Trim(sts[0],"\""),strings.Trim(sts[5],"\""),strings.Trim(sts[8],"\""))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Getfile() []Fileinfo {
	qsql := "SELECT id,filename FROM fileinfo"
	r,err := Db.Query(qsql)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var data []Fileinfo
	for r.Next(){
		var (
			id int64
			filename string
		)
		if err := r.Scan(&id,&filename); err != nil{
			fmt.Println(err)
			break
		}

		data = append(data,Fileinfo{
			Id:       id,
			Filename: filename,
		})

	}
	return data
}

func InitDB(dbtype,dsn string) error {
	var err error
	Db,err = sql.Open(dbtype,dsn)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err = Db.Ping();err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func CloseDb() error {
	return Db.Close()
}

func Readformdb(fileid int64) TotalData {
	ipsql := "select ipaddr,count(ipaddr) as total from recorder where fileid=? group by ipaddr order by total desc limit 10"
	codesql := "select status,count(status) from recorder where fileid=? group by status"
	var  tdata TotalData
	ir,err := Db.Query(ipsql,fileid)
	if err != nil {
		fmt.Println(err)
		return TotalData{}
	}
	tdata.Databar = procR(ir)
	cr,err := Db.Query(codesql,fileid)
	if err != nil {
		fmt.Println(err)
		return TotalData{}
	}
	tdata.Datapie = procR(cr)

	return tdata


}

func procR(r *sql.Rows) (data []map[string]string) {
	for r.Next(){
		tmap := make(map[string]string)
		var (
			name string
			value string
		)
		if err := r.Scan(&name,&value);err != nil {
			fmt.Println(err)
			return nil
		}
		tmap["name"] = name
		tmap["value"] = value
		data = append(data, tmap)
	}
	return data
}