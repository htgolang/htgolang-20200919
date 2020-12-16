package users

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"strings"
	"time"
)

var Db *sql.DB

type Userinfo struct {
	Id int 			`from:"-"`
	Name string		`from:"username"`
	Sex string		`from:"sex"`
	Addr string		`from:"addr"`
	Tel string		`from:"tel"`
	Birthday time.Time	`from:"brth"`
	Passwd string	`from:"password"`
}



//通过ID查找用户
func (u *Userinfo) FindByid(id int) (userinfo Userinfo,err error) {
	fsql := "SELECT  id, name, sex, addr,tel, birthday,passwd FROM users WHERE id = ? "
	var (
		uid int
		name string
		sex string
		addr string
		tel string
		passwd string
		birthday time.Time
	)
	err = Db.QueryRow(fsql,id).Scan(&uid,&name,&sex,&addr,&tel,&birthday,&passwd)
	if err != nil {
		return Userinfo{}, fmt.Errorf("未找到")
	}
	return Userinfo{
		Id:       uid,
		Name:     name,
		Sex:      sex,
		Addr:     addr,
		Tel:      tel,
		Birthday: birthday,
		Passwd:   passwd,
	},nil

}

// 新增用户
func (u *Userinfo) Add(name,addr,sex,tel,passwd,birthday string  )  error {
	var tname string
	qsql := "select name from users where name=?"

	_ = Db.QueryRow(qsql,name).Scan(&tname)
	if tname == name {
		return fmt.Errorf("%s 已存在.\n",name)
	}
	isql := "insert into users(name, sex, addr,tel,birthday,passwd,created_at, updated_at) values(?, ?, ?, ?, ?, ?, now(), now())"
	r,err := Db.Exec(isql,name,sex,addr,tel,birthday,passwd)
	//b,err:= time.ParseInLocation("2006-01-02", birthday,time.Local)
	if err != nil {
		return err
	}else {
		i,_ := r.LastInsertId()
		log.Printf("新增用户%s,Id:%d.\n",name,i)
	}

	return nil
}

//通过ID删除用户
func (u *Userinfo) Del(id int ) error {
	_,err := u.FindByid(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	dsql := "DELETE FROM users WHERE id = ?"
	r,err := Db.Exec(dsql,id)
	if err != nil {
		return err
	}
	_,_ = r.LastInsertId()
	log.Printf("ID:%d 已删除.\n",id)
	return nil

}
// 修改索引为index的用户信息
func (u *Userinfo) Modify(id,name,addr,sex,tel,bri,passwd string ) error {
	mid,err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("非法ID:%d\n",id)
	}
	for _,v := range u.Getall(){
		if mid == v.Id {
			continue
		}
		if v.Name == name {
			return fmt.Errorf("用户名:%s已存在.\n",name)
		}
	}

	usql := "UPDATE users SET name=?, addr=?, sex=?, tel=?, birthday=?, passwd=?  WHERE id=?"
	_,err = Db.Exec(usql,name,addr,sex,tel,bri,passwd,id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//通过关键字查找用户 返回userinfo结构体和error
func (u *Userinfo) QueryUser(str string) (f Userinfo, ok bool) {
	for _,v := range u.Getall() {
		if strings.Contains(v.Name,str) || strings.Contains(v.Addr,str)|| strings.Contains(v.Tel,str) {
			return v,true
		}
	}
	return Userinfo{}, false
}

//认证功能
func (u *Userinfo) Auth(username, passwd string) bool {
	//for _,v := range u.UserSlice{
	//	if username == v.Name && userutils.Summd5(passwd) == v.Passwd {
	//		return true
	//	}
	//}
	return false
}
//Get all
func (u *Userinfo) Getall() []Userinfo  {
	var userslice  []Userinfo
	qsql := "SELECT id,name,sex,addr,tel,birthday FROM users"
	r,err := Db.Query(qsql)
	if err != nil {
		fmt.Println(err)
	}
	for r.Next(){
		//var u Userinfo
		var (
			id int
			name string
			sex string
			addr string
			tel string
			birthday time.Time
		)
		err := r.Scan(&id,&name,&sex,&addr,&tel,&birthday)
		if err !=nil{
			fmt.Println(err)
		}
		u:= Userinfo{
			Id:       id,
			Name:     name,
			Sex:      sex,
			Addr:     addr,
			Tel:      tel,
			Birthday: birthday,
			Passwd:   "",
		}
		userslice = append(userslice,u)
	}
	if len(userslice) == 0 {
		return nil
	}else {
		return userslice
	}

}


func (u *Userinfo) InitDB(dbtype,dsn string) error {
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
func (u *Userinfo) CloseDb() error {
	return Db.Close()
}