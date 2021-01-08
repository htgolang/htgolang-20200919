package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

type User struct {
	ID        int64  `orm:"column(id);pk;auto"`
	Name      string `orm:"size(32)"`
	Password  string `orm:"size(1024)"`
	Sex       bool
	Tel       string     `orm:"size(32)"`
	Addr      string     `orm:"size(1024)"`
	CreatedAt *time.Time `orm:"auto_now_add;"`
	UpdatedAt *time.Time `orm:"auto_now;"`
	DeletedAt *time.Time `orm:"null;"`
	SortID    int64      `orm:"column(sort_id)"`
}

func NewUser(id int64, name string, sex bool, addr string, sort_id int64) *User {
	return &User{
		ID:       id,
		Name:     name,
		Password: "",
		Tel:      "",
		Sex:      sex,
		Addr:     addr,
		SortID:   sort_id,
	}
}

func init() {
	// 获取配置
	dsn := beego.AppConfig.String("db::dsN")

	// 注册驱动
	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		log.Fatal(err)
	}

	// 注册数据库
	if err := orm.RegisterDataBase("default", "mysql", dsn); err != nil {
		log.Fatal(err)
	}

	// 注册模型
	orm.RegisterModel(new(User))

	//orm.Debug = true
}
