package main

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 名称
// 驼峰式转化为下划线式 非首字母的大写字母前加_并将所有字符转化为小写
// 类型
// int integer
// int64 bigint
// float32 float
// float64 double
// string varchar(255)
// time.Time datetime
// bool boolean
// 修饰
// NOT NULL
// string bool int 默认值
// Id int* 并且未指定主键 自动将Id设置为自动增长的主键

// 标签 orm
// 列名 column
// 类型 type
// 		decimal类型 digits, decimals
// 字符串长度 size
// 主键 自动增长 pk auto
// 默认值 default
// 注释 description
// 允许为null  null
// 时间类型 auto_now_add; auto_now
// 索引 index, unique
type User struct { //
	Id        int64  `orm:"column(uid);pk;auto"`
	Name      string `orm:"size(64);default(aaaa)"`
	Password  string
	Sex       bool
	Height    float64    `orm:"digits(12);decimals(3)"`
	Tel       string     `orm:"index"`
	Addr      string     `orm:"type(text)"`
	Birthday  *time.Time `orm:"type(date)"`
	CreatedAt *time.Time `orm:"auto_now_add;description(创建时间)"`
	UpdatedAt *time.Time `orm:"auto_now"`
	DeletedAt *time.Time `orm:"null"`
}

func (u *User) TableName() string {
	return "a1"
}

func (u *User) TableIndex() [][]string {
	return [][]string{
		{"Name"},
		{"Password"},
		{"Tel", "Addr"},
	}
}

func main() {
	dsn := "golang:golang@2020@tcp(10.0.0.2:3306)/testorm?parseTime=true&loc=Local&charset=utf8mb4"

	// 导入驱动(初始化)
	// 导入orm包
	// 在ORM包中注册驱动(mysql)
	orm.RegisterDriver("mysql", orm.DRMySQL) // 可省略

	// 注册数据库(数据库的配置信息)
	if err := orm.RegisterDataBase("default", "mysql", dsn); err != nil {
		log.Fatal(err)
	}

	// 注册模型
	orm.RegisterModel(&User{}) // 指针类型的实例

	// 操作

	// DDL

	// 结构体对应表是否存在
	// 表不存在 创建对应的表
	// 若表存在 属性列是否在表中存在
	// 属性不存在 添加列
	// 索引是否存在 索引不存在 添加索引
	// 数据库别名
	// 是否先删除所有表
	// 显示详细信息
	orm.RunSyncdb("default", true, true) // 同步数据库
}
