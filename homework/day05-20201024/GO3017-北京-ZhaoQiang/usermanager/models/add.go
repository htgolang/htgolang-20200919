package models

func getID() int {
	id := 0
	for _, user := range users {
		if user.id > id {
			id = user.id
		}
	}
	return id
}

// AddUser 添加用户
func AddUser(u User) {
	u.id = getID() + 1
	users = append(users, u)
}
