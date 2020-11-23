package models

import (
	"strconv"
	"strings"
)

var users = []map[string]string{}

func getId() string {
	id := 0
	for _, user := range users {
		if uid, err := strconv.Atoi(user["id"]); err == nil {
			if uid > id {
				id = uid
			}
		}
	}
	return strconv.Itoa(id + 1)
}

func AddUser(user map[string]string) {
	user["id"] = getId()
	users = append(users, user)
}

func FindUserById(id string) map[string]string {
	for _, user := range users {
		if user["id"] == id {
			return user
		}
	}
	return nil
}
func ModifyUserById(user map[string]string, id string) {
	for idx, tuser := range users {
		if tuser["id"] == id {
			users[idx] = user
			break
		}
	}
}

func DeleteUserById(id string) {
	tempUsers := make([]map[string]string, 0, len(users)-1)
	for _, user := range users {
		if user["id"] != id {
			tempUsers = append(tempUsers, user)
		}
	}
	users = tempUsers
}

func filter(user map[string]string, q string) bool {
	return strings.Contains(user["name"], q) ||
		strings.Contains(user["addr"], q) ||
		strings.Contains(user["tel"], q)
}

func QueryUser(q string) []map[string]string {
	rt := make([]map[string]string, 0, len(users))
	for _, user := range users {
		if filter(user, q) {
			rt = append(rt, user)
		}
	}
	return rt
}

func ReturnUsers() []map[string]string {
	return users
}
