package main

import (
	"fmt"
	"sort"
)

type User struct {
	Id       int
	Name     string
	Password string
}

func (user User) String() string {
	return fmt.Sprintf("User[ID=%d, Name=%s]", user.Id, user.Name)
}

type Users []User

func (users Users) Len() int {
	return len(users)
}

func (users Users) Less(i, j int) bool {
	return users[i].Id < users[j].Id
}

func (users Users) Swap(i, j int) {
	users[i], users[j] = users[j], users[i]
}

func main() {
	var users Users = Users([]User{
		{10, "kk1", "password1"},
		{9, "kk3", "password3"},
		{7, "kk4", "password4"},
		{100, "kk5", "password5"},
		{88, "kk6", "password6"},
		{99, "kk7", "password7"},
	})
	var sl1 = []int{0, 33, 1, 3, 5, 3, 3, 7, 2}
	sort.Ints(sl1)
	sort.Sort(users)

	fmt.Println(users)
	fmt.Println(sl1)
	fmt.Println(sort.IsSorted(users))
	fmt.Println(sort.IntsAreSorted(sl1))
}
