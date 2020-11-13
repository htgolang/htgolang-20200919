package utils

import "sort"

func GetMaxId() int {
	if len(UsersList) > 0 {
		sort.Sort(UsersList)
		return UsersList[len(UsersList) - 1].Id
	}
	return 0
}