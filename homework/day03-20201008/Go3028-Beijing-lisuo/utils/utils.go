package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"text/tabwriter"

	define "github.com/htgolang/htgolang-20200919/tree/master/homework/day03-20201008/Go3028-Beijing-lisuo/define"
)

// gen a int64 number got Maximum 12 digits
func GenId() (res int64) {
	// gen a random number in [0, 999999999999)
	result, _ := rand.Int(rand.Reader, big.NewInt(999999999999))
	return result.Int64()
}

// to verify if a string contains only digits
func JustDigits(s string) bool {
	var a bool = true
	for _, c := range s {
		if c < '0' || c > '9' {
			a = false
			break
		}
	}
	return a
}

// show a user based on Id
func ShowUser(Id int64) {
	for _, userMap := range define.UserList {
		if val, ok := userMap[Id]; ok {
			//fmt.Println(i, Id, val)
			//fmt.Println(val.Name, val.Address, val.Phone)
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0|tabwriter.Debug)
			s := strconv.FormatInt(Id, 10)
			fmt.Fprintln(w, "|"+s+"\t"+val.Name+"\t"+val.Phone+"\t"+val.Address+" |")
			w.Flush()
		}
	}
}

// find user based on Id
func IdFindUser(UserList []map[int64]define.User, Id int64) define.User {
	var user define.User
	for _, userMap := range UserList {
		if _, ok := userMap[Id]; ok {
			user = userMap[Id]
		}
	}
	return user
}

// find user based on Name
func NameFindUser(Name string) map[int64]define.User {
	var user map[int64]define.User
	var id int64
	for _, userMap := range define.UserList {
		for k, v := range userMap {
			if v.Name == Name {
				user = map[int64]define.User{k: define.User{v.Name, v.Address, v.Phone}}
				id = k
			}
		}
	}
	if string(id) == "" {
		fmt.Println("No such user.")
	}
	return user
}

// del user based on Id
func IdDelUser(user *[]map[int64]define.User, id int64) {
	for i, u := range *user {
		for k, _ := range u {
			if int64(k) == id {
				*user = append(define.UserList[:i], define.UserList[i+1:]...)
			}
		}
	}
}

// del user based on Name
func NameDelUser(user *[]map[int64]define.User, name string) {
	for i, u := range *user {
		for _, v := range u {
			if v.Name == name {
				*user = append(define.UserList[:i], define.UserList[i+1:]...)
			}
		}
	}

}

// modify user based on Id
func IdModUser(user *[]map[int64]define.User, id int64) define.User {
	var input string
	var newUser define.User
	for _, u := range *user {
		for k, v := range u {
			if int64(k) == id {
				fmt.Printf("Input new Name [%v]: ", v.Name)
				fmt.Scanln(&input)
				if input == "" {
					newUser.Name = v.Name
				}
				newUser.Name = input
				fmt.Printf("Input new Address [%v]: ", v.Address)
				fmt.Scanln(&input)
				if input == "" {
					newUser.Address = v.Address
				}
				newUser.Address = input
				fmt.Printf("Input new Phone [%v]: ", v.Phone)
				fmt.Scanln(&input)
				if input == "" {
					newUser.Phone = v.Phone
				}
				newUser.Phone = input
				u[k] = newUser
				fmt.Printf("Modified user is: %v:%v\n", k, newUser)
			}
		}
	}
	return newUser
}

// modify user based on Name
func NameModUser(user *[]map[int64]define.User, name string) {
	var input string
	var newUser define.User
	for _, u := range *user {
		for k, v := range u {
			if v.Name == name {
				fmt.Printf("Input new Name [%v]: ", v.Name)
				fmt.Scanln(&input)
				if input == "" {
					newUser.Name = v.Name
				}
				newUser.Name = input
				fmt.Printf("Input new Address [%v]: ", v.Address)
				fmt.Scanln(&input)
				if input == "" {
					newUser.Address = v.Address
				}
				newUser.Address = input
				fmt.Printf("Input new Phone [%v]: ", v.Phone)
				fmt.Scanln(&input)
				if input == "" {
					newUser.Phone = v.Phone
				}
				newUser.Phone = input
				u[k] = newUser
				fmt.Printf("Modified user is: %v:%v\n", k, newUser)
			}
		}
	}
}
