package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	ID        int64      `orm:"column(id);pk;auto" json:"id"`
	Name      string     `orm:"size(32)" json:"name"`
	Password  string     `orm:"size(1024)" json:"-" xml:"-"`
	Sex       bool       `orm:"-" json:"sex"`
	Tel       string     `orm:"size(32)" json:"tel"`
	Addr      string     `orm:"size(1024)" json:"addr"`
	CreatedAt *time.Time `orm:"auto_now_add;" json:"created_at"`
	UpdatedAt *time.Time `orm:"auto_now;" json:"updated_at"`
	DeletedAt *time.Time `orm:"null;" json:"-"`
}

func NewUser(id int64, name string, sex bool, addr string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Password: "",
		Sex:      sex,
		Addr:     addr,
	}
}

func init() {
	orm.RegisterModel(new(User)) //&User{}
}
