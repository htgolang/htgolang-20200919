package users


// 定义用于存储user的db接口
type Mydb interface {
	FindByid(id int) (userinfo Userinfo,err error)
	Add(name,addr,sex,tel,passwd,birthday string  )  error
	Del(id int ) error
	Modify(id, name,addr,sex,tel,bri,passwd string ) error
	QueryUser(str string) (f Userinfo, ok bool)
	Auth(username, passwd string) bool
	Getall() []Userinfo
	InitDB(dbtype,dsn string) error
	CloseDb() error
}