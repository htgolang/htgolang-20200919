package users


// 定义用于存储user的db接口
type Mydb interface {
	Genid() int
	FindByid(id int) (userinfo Userinfo,index int,err error)
	Add(name,addr,tel,passwd,birthday string  )  error
	Del(id int ) error
	Modify(index int, name,addr,tel,bri string ) error
	QueryUser(str string) (f Userinfo, ok bool)
	Auth(username, passwd string) bool
	Load() error
	Sync() error
}