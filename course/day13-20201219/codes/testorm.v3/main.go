package main

import (
	"fmt"
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
	Id        int64      `orm:"column(uid);pk;auto"`
	Name      string     `orm:"size(64);"`
	Password  string     `orm:"size(1024);"`
	Tel       string     `orm:"size(32);"`
	Addr      string     `orm:"type(text)"`
	Birthday  *time.Time `orm:"type(date)"`
	CreatedAt *time.Time `orm:"auto_now_add"`
	UpdatedAt *time.Time `orm:"auto_now"`
	DeletedAt *time.Time `orm:"null"`
}

func main() {

	orm.Debug = true

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
	// orm.RunSyncdb("default", false, true) // 同步数据库

	//DML DQL

	ormer := orm.NewOrm()
	// // 对单个数据 增，删，改，查

	user := &User{Id: 10}
	fmt.Println(ormer.Read(user))
	fmt.Println(user)

	user = &User{Name: "kk"}

	fmt.Println(ormer.Read(user, "Name"))
	fmt.Println(user)

	user = &User{Name: "kk", Password: "123"}

	fmt.Println(ormer.Read(user, "Name", "Password"))
	fmt.Println(user)
}
