package model

//用户model定义

type User struct {
	Id      int
	Name    string
	Phone   string
	Address string
}

type UserDB struct {
	users []*User //用户列表
	count int     //用户数
	curId int     //当前用户的ID，用于用户ID自增
}

func NewUserDB() UserDB {
	return UserDB{
		users: []*User{},
		count: 0,
		curId: 0,
	}
}

func (u *UserDB) Create(name, phone, address string) error {
	u.curId++
	u.count++
	u.users = append(u.users, &User{
		Id:      u.curId,
		Name:    name,
		Phone:   phone,
		Address: address,
	})
	return nil
}

func (u *UserDB) List() ([]*User, error) {
	return u.users, nil
}
//todo 数据库方法