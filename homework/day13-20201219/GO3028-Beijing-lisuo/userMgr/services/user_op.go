package services

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"userMgr/models"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var (
	tableName = "user"
)

// Login handle user log logic
func Login() {
	fmt.Println("user logged in")
}

// ListAllUser list all users to home page
func ListAllUser() ([]*models.User, error) {
	var users []*models.User
	o := orm.NewOrm()
	qs := o.QueryTable(tableName)
	_, err := qs.All(&users)
	if err != nil {
		return []*models.User{}, err
	}
	return users, nil
}

// AddUser add a user to db
func CreateUser(name, password, address, cell string, sex int, born time.Time) error {
	o := orm.NewOrm()
	pass, errE := EncryptPass(password)
	if errE != nil {
		return errE
	}
	var user = models.User{
		Name:     name,
		Password: pass,
		Address:  address,
		Cell:     cell,
		Sex:      sex,
		Born:     &born,
	}
	_, err := o.Insert(&user)
	if err != nil {
		return err
	}
	return nil
}

// QueryUser get user from mysql based on id, name, address or cell
func QueryUser(id string, args ...string) ([]*models.User, error) {
	var userList []*models.User
	o := orm.NewOrm()
	qs := o.QueryTable(tableName)
	var tmpUsers = []*models.User{}
	if id != "" {
		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return userList, err
		}
		_, errq := qs.Filter("ID", Id).All(&tmpUsers)
		if errq != nil {
			return tmpUsers, errq
		}
		userList = append(userList, tmpUsers...)
		tmpUsers = nil
	} else {

		var blank = ""
		var i = 0
		var li = []string{"Name", "Address", "Cell"}
		fmt.Println("li: ", li)
		for _, v := range args {
			if v != blank {
				fmt.Println("filter: ", li[i]+"__icontains", v)
				_, errN := qs.Filter(li[i]+"__icontains", v).All(&tmpUsers)
				if errN != nil {
					return tmpUsers, errN
				}
				userList = append(userList, tmpUsers...)
				tmpUsers = nil
			}
			i++
		}
	}
	return userList, nil
}

// IDFindUser find user based on ID
func IDFindUser(Id int64) (models.User, error) {
	var user = models.User{}
	o := orm.NewOrm()
	user.ID = Id
	if err := o.Read(&user); err != nil {
		return user, err
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
func IDDelUser(id int64) error {
	o := orm.NewOrm()
	user := models.User{ID: id}
	_, err := o.Delete(&user)
	if err != nil {
		return err
	}
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

// IDModUser modify user based on ID
func IDModUser(name, address, password, cell, sex, born string, id int64) error {
	o := orm.NewOrm()
	s, erra := strconv.Atoi(sex)
	if erra != nil {
		return erra
	}
	b, errp := time.Parse("2006-01-02", born)
	if errp != nil {
		return errp
	}
	user := models.User{
		ID:       id,
		Name:     name,
		Address:  address,
		Password: password,
		Cell:     cell,
		Sex:      s,
		Born:     &b,
	}

	_, errU := o.Update(&user)
	if errU != nil {
		return errU
	}
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

// EncryptPass en
func EncryptPass(pass string) (string, error) {
	bt, err := bcrypt.GenerateFromPassword([]byte(pass), 8)
	return string(bt), err
}
