package services

import (
	"crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"userMgr/models"
	"userMgr/utils"

	_ "github.com/go-sql-driver/mysql"
)

// ListAllUser list all users to home page
func ListAllUser() (models.UserList, error) {
	var users models.UserList
	rows, err := models.DB.Query("SELECT * FROM user")
	if err != nil {
		fmt.Println(err)
		return nil, err
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
			user := models.NewUser(id, sex, name, cell, address, password, *born)
			users = append(users, user)
			//fmt.Println(id, name, password, sex, born, cell, created_at, updated_at, deleted_at)
		}
	}
	return users, nil
}

// QueryUser get user from mysql based on id, name, address or cell
func QueryUser(id, name, address, cell string) (models.UserList, error) {
	var userList = models.UserList{}
	var blank = ""
	f := func(args ...string) map[int]string {
		var queryStrs = make(map[int]string)
		for i, v := range args {
			if v != blank {
				queryStrs[i] = "%" + v + "%"
			} else {
				queryStrs[i] = v
			}
		}
		return queryStrs

	}(name, address, cell)
	fmt.Printf("query string: %#v\n", f)

	s := `SELECT * FROM user WHERE id = ? OR name LIKE ? OR address LIKE ? OR cell LIKE ?`
	rows, err := models.DB.Query(s, id, f[0], f[1], f[2])
	fmt.Println("query sql: ", s)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for rows.Next() {
		var (
			id         int64
			name       string
			sex        int
			address    string
			cell       string
			born       string
			passwd     string
			created_at string
			updated_at string
			deleted_at sql.NullString
		)
		if err := rows.Scan(&id, &name, &passwd, &sex, &born, &address, &cell, &created_at, &updated_at, &deleted_at); err != nil {
			fmt.Println(err)
			return nil, err
		}
		fmt.Println("query id and name: ", id, name)
		t, err := time.Parse("2006-01-02T00:00:00+08:00", born)
		if err != nil {
			return nil, err
		}
		user := models.User{
			ID:      id,
			Name:    name,
			Sex:     sex,
			Address: address,
			Cell:    cell,
			Born:    t,
			Passwd:  passwd,
		}
		userList = append(userList, user)
	}
	return userList, nil
}

// AddUser add a user to db
func CreateUser(name, password, address, cell string, sex int, born time.Time) error {
	sql := `
    INSERT INTO user (name, password, sex, born, address, cell, created_at, updated_at) values (?, password(?), ?, ?, ?, ?, NOW(), NOW()) `
	result, err := models.DB.Exec(sql, name, password, sex, born, address, cell)
	if err != nil {
		return err
	}
	fmt.Println("LastInsertID and RowsAffected: ")
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	return nil
}

// IDFindUser find user based on ID
func IDFindUser(Id int64) (models.User, error) {
	var user = models.User{}
	sql := `
    SELECT * FROM user WHERE id = ?
    `
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
	err := models.DB.QueryRow(sql, Id).Scan(&id, &name, &password, &sex, &born, &address, &cell, &created_at, &updated_at, &deleted_at)
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	} else {
		user = models.User{
			ID:      id,
			Name:    name,
			Sex:     sex,
			Address: address,
			Cell:    cell,
			Born:    *born,
			Passwd:  password,
		}
		fmt.Println("IDFind user: ")
		fmt.Println(id, name, password, sex, born, cell, created_at, updated_at, deleted_at)
	}
	return user, nil
}

// NameFindUser find user based on Name
func NameFindUser(Name string) error {
	sql := `
    SELECT * FROM user WHERE name = ?
    `
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
	err := models.DB.QueryRow(sql, Name).Scan(&id, &name, &password, &sex, &born, &address, &cell, &created_at, &updated_at, &deleted_at)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Println(id, name, password, sex, born, cell, created_at, updated_at, deleted_at)
	}
	return nil
}

// IDDelUser del user based on ID
func IDDelUser(ID int64) error {
	sql := `
    DELETE FROM user WHERE id = ?
    `
	result, err := models.DB.Exec(sql, ID)
	if err != nil {
		return err
	}
	fmt.Println("LastInsertId:")
	fmt.Println(result.LastInsertId())
	fmt.Println("RowsAffected:")
	fmt.Println(result.RowsAffected())
	return nil
}

// NameDelUser del user based on Name
func NameDelUser(db *sql.DB, Name string) error {
	sql := `
    DELETE FROM user WHERE name = ?
    `
	result, err := db.Exec(sql, Name)
	if err != nil {
		return err
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	return nil
}

// GetMaxID get max id of current models.UserList
func GetMaxID(db *sql.DB) (int64, error) {
	sql := `
    SELECT * FROM user ORDER BY id DESC LIMIT 1
	`
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
	err := db.QueryRow(sql).Scan(&id, &name, &password, &sex, &born, &address, &cell, &created_at, &updated_at, &deleted_at)
	if err != nil {
		return id, err
	}
	return id, nil
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
func IDModUser(name, address, password, cell, sex, born string, id int64) error {
	sql := `UPDATE user SET name= ?, sex= ?, password=password(?), born= ?, address= ?, cell= ? WHERE id= ?;`
	result, err := models.DB.Exec(sql, name, sex, password, born, address, cell, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("mod user %v LastInsertID and RowsAffected: ", id)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	return nil
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
