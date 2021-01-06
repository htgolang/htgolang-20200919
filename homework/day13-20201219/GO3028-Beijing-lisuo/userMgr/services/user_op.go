package services

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"userMgr/forms"
	"userMgr/models"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var (
	tableName = "user"
)

// Login handle user log logic
func LoginAuth(form *forms.AuthForm) (*models.User, error) {
	user, err := NameFindUser(form.UserName)
	fmt.Println("loginauth user, err", user, err)
	if err != nil {
		// user name wrong
		return nil, errors.New("user name wrong, no such user")
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.PassWord)); err != nil {
			// user pass wrong
			return nil, errors.New("user password wrong")
		} else {
			return user, nil
		}
	}
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
	_, errf := NameFindUser(name)
	if errf == nil {
		return errors.New("user exists, do not create same name user")
	}
	o := orm.NewOrm()
	pass, errE := EncryptPass(password, models.PassCost)
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
func NameFindUser(n string) (*models.User, error) {
	var user models.User
	user.Name = n
	o := orm.NewOrm()
	if err := o.Read(&user, "Name"); err != nil {
		return &user, err
		fmt.Println("name find user: ", &user, err)
	}
	return &user, nil
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
	p, err := EncryptPass(password, models.PassCost)
	if err != nil {
		return err
	}
	user := models.User{
		ID:       id,
		Name:     name,
		Address:  address,
		Password: p,
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

//IfAdmin detect if admin is exists when app run
// if not exists, create a admin
func IfAdmin() error {
	o := orm.NewOrm()
	b := time.Now()
	p, err := EncryptPass(models.AdminPass, models.PassCost)
	if err != nil {
		return err
	}
	var adminUser = models.User{
		ID:       models.AdminID,
		Name:     models.AdminName,
		Sex:      1,
		Address:  "Beijing",
		Password: p,
		Cell:     "18811738844",
		Born:     &b,
	}
	_, errf := NameFindUser(models.AdminName)
	if errf != nil {
		_, err := o.Insert(&adminUser)
		if err != nil {
			return err
		}
		return errf
	} else {
		fmt.Println("admin is exists.")
		return nil
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
func EncryptPass(pass string, cost int) (string, error) {
	bt, err := bcrypt.GenerateFromPassword([]byte(pass), models.PassCost)
	return string(bt), err
}
