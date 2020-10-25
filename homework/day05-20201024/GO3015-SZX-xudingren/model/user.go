package model

//用户model定义
type User struct {
	Id      int
	Name    string
	Phone   string
	Address string
}

//工厂函数
func NewUser(name, phone, address string) User {
	return User{
		Name:    name,
		Phone:   phone,
		Address: address,
	}
}
