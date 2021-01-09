package main

import (
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Role struct {
	Id   int64
	Name string

	Users []*User `orm:"reverse(many);"`
	Urls  []*Url  `orm:"rel(m2m)"`
}

type User struct {
	Id   int64
	Name string

	Password *Password `orm:"reverse(one)"`
	Role     *Role     `orm:"rel(fk);"`
}

type Password struct {
	Id   int64
	Hash string

	User *User `orm:"rel(one)"`
}

type Url struct {
	Id   int64
	Name string
	Path string

	Roles []*Role `orm:"reverse(many)"`
}

/*
自动生成关系表
type RoleUrl struct {
	Id     int64
	RoleId int64
	UrlId  int64
}
*/

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

	if err := orm.RegisterDataBase("default", "mysql", dsn); err != nil {
		log.Fatal(err)
	}

	orm.RegisterModel(&Role{}, &User{}, &Password{}, &Url{}) // 指针类型的实例

	orm.RunSyncdb("default", true, false) // 同步数据库

	orm.Debug = true
	ormer := orm.NewOrm()
	role := &Role{Name: "管理员"}
	ormer.Insert(role)

	user := &User{
		Name: "root",
		Role: role,
	}
	ormer.Insert(user)

	fmt.Println("insert")
	time.Sleep(10 * time.Second)
	ormer.Delete(role)

	/*
		password := &Password{
			Hash: "root@password",
			User: user,
		}
		ormer.Insert(password)

		u1 := &Url{
			Name: "新增用户",
			Path: "/user/add",
		}
		u2 := &Url{
			Name: "删除用户",
			Path: "/user/delete",
		}
		ormer.Insert(u1)
		ormer.Insert(u2)

		m2m := ormer.QueryM2M(role, "Urls")
		m2m.Add(u1, u2)
	*/
}
