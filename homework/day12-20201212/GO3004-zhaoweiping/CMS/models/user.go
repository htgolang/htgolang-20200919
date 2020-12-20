package models

type User struct {
	ID   int64
	Name string
	Sex  bool
	Addr string
}

func NewUser(id int64, name string, sex bool, addr string) *User {
	// fmt.Println(name)
	return &User{id, name, sex, addr}
}
