package config

import (
	"database/sql"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

// 定义数据库配置结构体
type DbConf struct {
	Dbip string `yaml:"ip"`
	Dbport string `yaml:"port"`
	Dbname string `yaml:"database"`
	Dbuser string `yaml:"user"`
	Dbpass string `yaml:"passwd"`
}

func NewDbConf() *DbConf {
	return &DbConf{}
}

// 定义从yaml文件读取数据库配置函数
func (c *DbConf) getConf() error {
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Printf("无法打开配置文件:%v\n", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Printf("解析yaml文件失败，请检查配置文件:%v\n", err)
		return err
	}
	return nil
}

// 获取数据库连接
func (c *DbConf) InitDb() (error) {
	// 获取配置
	err := c.getConf()
	if err != nil {
		return err
	}

	// 连接数据库
	driver := "mysql"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Local&charset=utf8mb4",
		c.Dbuser, c.Dbpass, c.Dbip, c.Dbport, c.Dbname)
	Db, err = sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	err = Db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func CloseDb() error {
	return Db.Close()
}