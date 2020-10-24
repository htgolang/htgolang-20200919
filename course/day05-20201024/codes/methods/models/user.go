package models

type User struct {
	id   int
	name string
	age  int
}

func NewUser(id int, name string, age int) User {
	return User{id, name, age}
}

// 不能修改调用者的age
func (user User) AddAge() {
	user.age += 1
}

func (user User) GetName() string {
	return user.name
}

// 不能修改函数实参中的age
func AddAge(user User) {
	user.age += 1
}

func GetName(user User) string {
	return user.name
}
