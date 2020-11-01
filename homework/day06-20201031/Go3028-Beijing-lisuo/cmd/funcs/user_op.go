package funcs

import (
	"crypto/md5"
	"errors"
	"fmt"
	"time"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day06-20201031/Go3028-Beijing-lisuo/cmd/db"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day06-20201031/Go3028-Beijing-lisuo/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day06-20201031/Go3028-Beijing-lisuo/utils"
)

// NewUser make a new user contains user's info
func NewUser(id int64, name, cell, address, born, passwd string) define.User {
	return define.User{
		ID:      id,
		Name:    name,
		Cell:    cell,
		Address: address,
		Born: func() time.Time {
			t, _ := time.Parse("2006.01.02", born)
			return t
		}(),
		Passwd: md5.Sum([]byte(passwd)),
	}
}

// Init add some users to define.UserList
func Init(ul *[]define.User) {
	user0 := NewUser(0, "admin", "18811992299", "HaidianDistrict,BeijingXinParkRestaurants,BeixiaguanSubdistrict,HaidianDistrict,China",
		time.Now().Format("2006.01.02"), "qwert")
	//user1 := NewUser(1, "jack ma", "18800009999", "Hangzhou, China", time.Now().Format("2006.01.02"), "hello")
	//user3 := NewUser(3, "steve", "18800002222", "Mars", time.Now().Format("2006.01.02"), "hi")
	(*ul) = append((*ul), user0)
	//fmt.Printf("user %v added\n", user0.Name)
	//fmt.Printf("user %v added\n", user1.Name)
	//fmt.Printf("user %v added\n", user3.Name)
	AddFunc()
	db.ReadUsers()
}

// IDFindUser find user based on ID
func IDFindUser(ul *[]define.User, id int64) (define.User, error) {
	for _, user := range *ul {
		if user.ID == id {
			return user, nil
		}
	}
	err := errors.New("no such user")
	return define.User{}, err
}

// NameFindUser find user based on Name
func NameFindUser(ul *[]define.User, Name string) (define.User, error) {
	for _, user := range *ul {
		if user.Name == Name {
			return user, nil
		}
	}
	err := errors.New("no such user")
	return define.User{}, err
}

// IDDelUser del user based on ID
func IDDelUser(ul *[]define.User, id int64) error {
	for i, user := range *ul {
		if int64(user.ID) == id {
			if id == define.AdminID {
				err := errors.New("you'r not allowed to modify admin, nothing changed")
				return err
			}
			*ul = append(define.UserList[:i], define.UserList[i+1:]...)
		}
	}
	return nil
}

// NameDelUser del user based on Name
func NameDelUser(ul *[]define.User, name string) error {
	var index int
	idx, err := GetUserIndex(ul, name)
	if err != nil {
		fmt.Println(err)
	}
	if name == define.AdminName {
		err := errors.New("you'r not allowed to modify admin, nothing changed")
		return err
	} else if (*ul)[idx].Name == name {
		index = idx
	}
	(*ul) = append(define.UserList[:index], define.UserList[index+1:]...)
	return nil
}

// GetMaxID get max id of current define.UserList
func GetMaxID(ul *[]define.User) int64 {
	var MaxID int64 = -1
	for _, user := range *ul {
		var i int64 = user.ID
		if i > MaxID {
			MaxID = i
		}
	}
	return MaxID
}

// GetUserIndex return a user's index in define.UserList and a error
func GetUserIndex(ul *[]define.User, name string) (int, error) {
	var index int = -1
	for i, u := range *ul {
		if u.Name == name {
			index = i
			return index, nil
		}
	}
	return index, errors.New("not fund index")
}

// IDModUser modify user based on ID
func IDModUser(ul *[]define.User, id int64) (define.User, error) {
	var iname, iaddress, cell, ipasswd string
	newUser := define.User{ID: id}
	if id == define.AdminID {
		err := errors.New("you'r not allowed to modify admin, nothing changed")
		return newUser, err
	}
	for _, user := range *ul {
		if user.ID == id {
			fmt.Println("modifying...........")
			fmt.Printf("Input new Name [%v]: ", user.Name)
			iname = utils.Read()
			newUser.Name = iname
			if iname == "" {
				newUser.Name = user.Name
			}

			fmt.Printf("Input new Address [%v]: ", user.Address)
			iaddress = utils.Read()
			newUser.Address = iaddress
			if iaddress == "" {
				newUser.Address = user.Address
			}

			fmt.Printf("Input new Phone [%v]: ", user.Cell)
			cell = utils.Read()
			// make sure the phone number contains only pure digits
			for utils.JustDigits(cell) == false {
				fmt.Print("Please input a legal phone number: \n> ")
				cell = utils.Read()
				if utils.JustDigits(cell) == true {
					break
				}
			}
			newUser.Cell = cell
			if cell == "" {
				newUser.Cell = user.Cell
			}

			fmt.Printf("Input new passwd [%v]: ", user.Passwd)
			ipasswd = utils.Read()
			newUser.Passwd = md5.Sum([]byte(ipasswd))
			if ipasswd == "" {
				newUser.Passwd = user.Passwd
			}
		}
	}
	return newUser, nil
}

// NameModUser modify user based on Name
func NameModUser(ul *[]define.User, name string) (define.User, error) {
	var iname, iaddress, iphone, ipasswd string
	//var index int
	newUser := define.User{}
	if name == define.AdminName {
		err := errors.New("you'r not allowed to modify admin, nothing changed")
		return newUser, err
	}
	for _, u := range *ul {
		if u.Name == name {

			fmt.Println("modifying...........")
			newUser.ID = u.ID
			fmt.Printf("Input new Name [%v]: ", u.Name)
			iname = utils.Read()
			newUser.Name = iname
			if iname == "" {
				newUser.Name = u.Name
			}

			fmt.Printf("Input new Address [%v]: ", u.Address)
			iaddress = utils.Read()
			newUser.Address = iaddress
			if iaddress == "" {
				newUser.Address = u.Address
			}

			fmt.Printf("Input new Phone [%v]: ", u.Cell)
			iphone = utils.Read()
			// make sure the phone number contains only pure digits
			for utils.JustDigits(iphone) == false {
				fmt.Print("Please input a legal phone number: \n> ")
				iphone = utils.Read()
				if utils.JustDigits(iphone) == true {
					break
				}
			}
			newUser.Cell = iphone
			if iphone == "" {
				newUser.Cell = u.Cell
			}

			fmt.Printf("Input new passwd [%v]: ", u.Passwd)
			ipasswd = utils.Read()
			newUser.Passwd = md5.Sum([]byte(ipasswd))
			if ipasswd == "" {
				newUser.Passwd = u.Passwd
			}
			fmt.Printf("Modified user is: %v:%v\n", newUser.Name, newUser)
			return newUser, nil
		}
	}
	return define.User{}, nil
}
