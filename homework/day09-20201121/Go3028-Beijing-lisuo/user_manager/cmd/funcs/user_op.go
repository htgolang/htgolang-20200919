package funcs

import (
	"crypto/md5"
	"errors"
	"fmt"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day09-20201121/Go3028-Beijing-lisuo/user_manager/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day09-20201121/Go3028-Beijing-lisuo/user_manager/utils"
)

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
	var index int
	newUser := define.User{ID: id}
	if id == define.AdminID {
		err := errors.New("you'r not allowed to modify admin, nothing changed")
		return newUser, err
	}
	for i, user := range *ul {
		if user.ID == id {
			index = i
			fmt.Println("modifying...........")
			fmt.Printf("Input new Name [%v]: ", user.Name)
			iname = utils.Read()
			if iname != "" {
				newUser.Name = iname
			} else {
				newUser.Name = user.Name
			}

			fmt.Printf("Input new Address [%v]: ", user.Address)
			iaddress = utils.Read()
			if iaddress != "" {
				newUser.Address = iaddress
			} else {
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
			if cell != "" {
				newUser.Cell = cell
			} else {
				newUser.Cell = user.Cell
			}

			fmt.Printf("Input new passwd [%v]: ", user.Passwd)
			ipasswd = fmt.Sprintf("%x", md5.Sum([]byte(utils.Read())))
			if ipasswd != "" {
				newUser.Passwd = ipasswd
			} else {
				newUser.Passwd = user.Passwd
			}
			newUser.Born = user.Born
		}

	}
	(*ul)[index] = newUser
	return (*ul)[index], nil
}

// NameModUser modify user based on Name
func NameModUser(ul *[]define.User, name string) (define.User, error) {
	var iname, iaddress, iphone, ipasswd string
	var index int
	newUser := define.User{}
	if name == define.AdminName {
		fmt.Println("in if")
		return newUser,
			errors.New("you'r not allowed to modify admin, nothing changed")
	}
	for i, u := range *ul {
		if u.Name == name {
			fmt.Printf("range if u.Name: %v, name: %v\n", u.Name, name)

			index = i
			fmt.Println("modifying...........")
			newUser.ID = u.ID
			fmt.Printf("Input new Name [%v]: ", u.Name)
			iname = utils.Read()
			if iname != "" {
				newUser.Name = iname
			} else {
				newUser.Name = u.Name
			}

			fmt.Printf("Input new Address [%v]: ", u.Address)
			iaddress = utils.Read()
			if iaddress != "" {
				newUser.Address = iaddress
			} else {
				newUser.Address = u.Address

			}

			fmt.Printf("Input new Phone [%v]: ", u.Cell)
			iphone = utils.Read()
			if iphone != "" {
				// make sure the phone number contains only pure digits
				for utils.JustDigits(iphone) == false {
					fmt.Print("Please input a legal phone number: \n> ")
					iphone = utils.Read()
					if utils.JustDigits(iphone) == true {
						break
					}
				}
				newUser.Cell = iphone
			} else {
				newUser.Cell = u.Cell
			}

			fmt.Printf("Input new passwd [%v]: ", u.Passwd)
			ipasswd = fmt.Sprintf("%x", md5.Sum([]byte(utils.Read())))
			if ipasswd != "" {
				newUser.Passwd = ipasswd
			} else {
				newUser.Passwd = u.Passwd
			}
			newUser.Born = u.Born
			(*ul)[index] = newUser
			return (*ul)[index], nil
		}
	}
	return newUser, errors.New("not modified")
}
