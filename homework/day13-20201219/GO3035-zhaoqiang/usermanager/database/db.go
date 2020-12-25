package database

import (
	"fmt"
	"os"
	"time"
	"user/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbuser := beego.AppConfig.DefaultString("db::user", "")
	dbpass := beego.AppConfig.DefaultString("db::password", "")
	host := beego.AppConfig.DefaultString("db::host", "localhost")
	port := beego.AppConfig.DefaultString("db::port", "3306")
	database := beego.AppConfig.String("db::database")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local&parseTime=true", dbuser, dbpass, host, port, database)
	driverName := "mysql"
	err := orm.RegisterDriver(driverName, orm.DRMySQL)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	err = orm.RegisterDataBase("default", driverName, dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	orm.RegisterModel(new(models.User))
	orm.RunSyncdb("default", false, true)
	ormer := orm.NewOrm()

	user := &models.User{ID: 1}
	err = ormer.Read(user)
	if err != nil {
		sql := `insert into
		userorm(id, name, password, admin, brithday, create_at, update_at)
		values(1, 'admin', '$2a$05$9c8VHzyyJawQXHn55gEdxOOfQjxX.D0DKOEar1.kdrnUGeR6qIu8S', 1, ?, ?, ?)`
		rawseter := ormer.Raw(sql, time.Now(), time.Now(), time.Now())
		_, err = rawseter.Exec()
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}
}
