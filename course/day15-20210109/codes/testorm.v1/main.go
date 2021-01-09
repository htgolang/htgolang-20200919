package main

import (
	"log"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Role struct {
	Id   int64
	Name string
}

type User struct {
	Id     int64
	Name   string
	RoleId int64
}

func (u *User) TableName() string {
	return "user2"
}

type Password struct {
	Id     int64
	Hash   string
	UserId int64
}

type Url struct {
	Id   int64
	Name string
	Path string
}

type RoleUrl struct {
	Id     int64
	RoleId int64
	UrlId  int64
}

// User Password 1:1
// User Role 1:n
// Role User n:m

// 主表删除 -> 附属表 数据
// 		删除(常用)
// 		设置为NULL
// 		不动
// 		设置为默认值

// 1:1
// User Password
// orm: Password -> User rel(one)
// orm: User Password reverse(one)

// 1:n
// User Role
// orm: User -> Role rel(fk)
// orm: Role -> []User reverse(many)

// m:n

// Role Url
// orm: Role -> []URl rel(m2m)
// orm： Url -> []Role reverse(many)
func main() {
	dsn := "golang:golang@2020@tcp(10.0.0.2:3306)/testorm?parseTime=true&loc=Local&charset=utf8mb4"
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
