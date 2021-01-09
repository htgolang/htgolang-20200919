package main

import (
	"fmt"
	"log"

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

	// r2 := &Role{Id: 1}

	// ormer.Read(r2)
	// ormer.LoadRelated(r2, "Users")
	// fmt.Println(r2.Users)

	// u2 := &User{Id: 1}
	// ormer.Read(u2)
	// ormer.LoadRelated(u2, "Role")
	// fmt.Println(u2.Role)

	password := &Password{
		Hash: "root@password",
		User: user,
	}
	ormer.Insert(password)

	// u2 := &User{Id: 1}
	// ormer.Read(u2)
	// ormer.LoadRelated(u2, "Password")
	// fmt.Println(u2.Password)

	// p2 := &Password{Id: 1}
	// ormer.Read(p2)
	// ormer.LoadRelated(p2, "User")
	// fmt.Println(p2.User)

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
	// time.Sleep(10 * time.Second)

	// m2m.Remove(u2)
	// m2m.Clear()
	// fmt.Println(m2m.Count())
	// fmt.Println(m2m.Exist(u2))

	/*
		r2 := &Role{Id: 1}
		ormer.Read(r2)
		ormer.LoadRelated(r2, "Urls")
		fmt.Println(r2.Urls)

		uu2 := &Url{Id: 1}
		ormer.Read(uu2)
		ormer.LoadRelated(uu2, "Roles")
		fmt.Println(uu2.Roles)
	*/

	queryset := ormer.QueryTable(&User{})
	var users []*User
	queryset.RelatedSel().Filter("Password__Hash__icontains", "a").All(&users)
	fmt.Println(users[0].Role)
	fmt.Println(users[0].Password)

	queryset = ormer.QueryTable(&Role{})
	var roles []*Role
	queryset.RelatedSel().Filter("Urls__Url__Path__icontains", "user").All(&roles)
	fmt.Println(roles)
}
