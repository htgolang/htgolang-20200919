package services

import (
	"crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"user_manager/models"
	"user_manager/user_utils"

	_ "github.com/go-sql-driver/mysql"
)

func ListAllUser(db *sql.DB) error {
	rows, err := db.Query("select * from user")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id         int64
			name       string
			password   string
			sex        int
			born       *time.Time
			address    string
			cell       string
			created_at *time.Time
			updated_at *time.Time
			deleted_at *time.Time
		)
		err := rows.Scan(&id, &name, &password, &sex, &born, &address, &cell, &created_at, &updated_at, &deleted_at)
		if err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(id, name, password, sex, born, cell, created_at, updated_at, deleted_at)
		}
	}
	return nil
}

// IDFindUser find user based on ID
func IDFindUser(ul *[]models.User, id int64) (models.User, error) {
	for _, user := range *ul {
		if user.ID == id {
			return user, nil
		}
	}
	err := errors.New("no such user")
	return models.User{}, err
}

// NameFindUser find user based on Name
func NameFindUser(ul *[]models.User, Name string) (models.User, error) {
	for _, user := range *ul {
		if user.Name == Name {
			return user, nil
		}
	}
	err := errors.New("no such user")
	return models.User{}, err
}

// IDDelUser del user based on ID
func IDDelUser(ul *[]models.User, id int64) error {
	for i, user := range *ul {
		if int64(user.ID) == id {
			if id == models.AdminID {
				err := errors.New("you'r not allowed to modify admin, nothing changed")
				return err
			}
			*ul = append(models.UserList[:i], models.UserList[i+1:]...)
		}
	}
	return nil
}

// NameDelUser del user based on Name
func NameDelUser(ul *[]models.User, name string) error {
	var index int
	idx, err := GetUserIndex(ul, name)
	if err != nil {
		fmt.Println(err)
	}
	if name == models.AdminName {
		err := errors.New("you'r not allowed to modify admin, nothing changed")
		return err
	} else if (*ul)[idx].Name == name {
		index = idx
	}
	(*ul) = append(models.UserList[:index], models.UserList[index+1:]...)
	return nil
}

// GetMaxID get max id of current models.UserList
func GetMaxID(ul *[]models.User) int64 {
	var MaxID int64 = -1
	for _, user := range *ul {
		var i int64 = user.ID
		if i > MaxID {
			MaxID = i
		}
	}
	return MaxID
}

// GetUserIndex return a user's index in models.UserList and a error
func GetUserIndex(ul *[]models.User, name string) (int, error) {
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
func IDModUser(ul *[]models.User, id int64) (models.User, error) {
	var iname, iaddress, cell, ipasswd string
	var index int
	newUser := models.User{ID: id}
	if id == models.AdminID {
		err := errors.New("you'r not allowed to modify admin, nothing changed")
		return newUser, err
	}
	for i, user := range *ul {
		if user.ID == id {
			index = i
			fmt.Println("modifying...........")
			fmt.Printf("Input new Name [%v]: ", user.Name)
			iname = user_utils.Read()
			if iname != "" {
				newUser.Name = iname
			} else {
				newUser.Name = user.Name
			}

			fmt.Printf("Input new Address [%v]: ", user.Address)
			iaddress = user_utils.Read()
			if iaddress != "" {
				newUser.Address = iaddress
			} else {
				newUser.Address = user.Address
			}

			fmt.Printf("Input new Phone [%v]: ", user.Cell)
			cell = user_utils.Read()
			// make sure the phone number contains only pure digits
			for user_utils.JustDigits(cell) == false {
				fmt.Print("Please input a legal phone number: \n> ")
				cell = user_utils.Read()
				if user_utils.JustDigits(cell) == true {
					break
				}
			}
			if cell != "" {
				newUser.Cell = cell
			} else {
				newUser.Cell = user.Cell
			}

			fmt.Printf("Input new passwd [%v]: ", user.Passwd)
			ipasswd = fmt.Sprintf("%x", md5.Sum([]byte(user_utils.Read())))
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
func NameModUser(ul *[]models.User, name string) (models.User, error) {
	var iname, iaddress, iphone, ipasswd string
	var index int
	newUser := models.User{}
	if name == models.AdminName {
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
			iname = user_utils.Read()
			if iname != "" {
				newUser.Name = iname
			} else {
				newUser.Name = u.Name
			}

			fmt.Printf("Input new Address [%v]: ", u.Address)
			iaddress = user_utils.Read()
			if iaddress != "" {
				newUser.Address = iaddress
			} else {
				newUser.Address = u.Address

			}

			fmt.Printf("Input new Phone [%v]: ", u.Cell)
			iphone = user_utils.Read()
			if iphone != "" {
				// make sure the phone number contains only pure digits
				for user_utils.JustDigits(iphone) == false {
					fmt.Print("Please input a legal phone number: \n> ")
					iphone = user_utils.Read()
					if user_utils.JustDigits(iphone) == true {
						break
					}
				}
				newUser.Cell = iphone
			} else {
				newUser.Cell = u.Cell
			}

			fmt.Printf("Input new passwd [%v]: ", u.Passwd)
			ipasswd = fmt.Sprintf("%x", md5.Sum([]byte(user_utils.Read())))
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
