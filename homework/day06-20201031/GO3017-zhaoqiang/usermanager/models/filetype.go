package models

import (
	"encoding/csv"
	"encoding/gob"
	"os"
	"strconv"
	"time"
	"zhao/utils"
)

type Files interface {
	Storage([]User) error
	LoadUsers() error
	InitFile() error
}

type GobPersistence struct {
	Path string
}

func (p GobPersistence) LoadUsers() error {
	file, err := os.Open(p.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	gob.Register(User{})
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil {
		return err
	}
	return nil
}

func (p GobPersistence) Storage(users []User) error {
	file, err := os.Create(p.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	gob.Register(User{})
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		return err
	}
	return nil
}

func (p GobPersistence) InitFile() error {
	passwdstr, err := utils.GetPasswd("init admin passwd:")
	if err != nil {
		return err
	}
	passwd := utils.Md5Convert([]byte(passwdstr))
	user := User{
		ID:       0,
		Name:     "admin",
		Addr:     "admin",
		Tel:      "admin",
		Birthday: time.Now(),
		Passwd:   passwd,
	}
	return p.Storage([]User{user})
}

//csv 格式
type CsvPersistence struct {
	Path string
}

func (p CsvPersistence) LoadUsers() error {
	tmpusers := make([]User, 0, 10)
	file, err := os.Open(p.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	recodes, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, line := range recodes {
		id, _ := strconv.Atoi(line[0])
		brithstr, _ := time.Parse("20060102", line[4])

		user := User{
			ID:       id,
			Name:     line[1],
			Addr:     line[2],
			Tel:      line[3],
			Birthday: brithstr,
			Passwd:   line[5],
		}
		tmpusers = append(tmpusers, user)
	}
	users = tmpusers
	return nil
}

func (p CsvPersistence) Storage(users []User) error {
	file, err := os.Create(p.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	csvw := csv.NewWriter(file)
	for _, user := range users {
		recode := []string{
			0: strconv.Itoa(user.ID),
			1: user.Name,
			2: user.Addr,
			3: user.Tel,
			4: user.Birthday.Format("20060102"),
			5: user.Passwd,
		}
		csvw.Write(recode)
	}
	csvw.Flush()
	return nil
}

func (p CsvPersistence) InitFile() error {
	passwdstr, err := utils.GetPasswd("init admin passwd:")
	if err != nil {
		return err
	}
	passwd := utils.Md5Convert([]byte(passwdstr))
	user := User{
		ID:       0,
		Name:     "admin",
		Addr:     "admin",
		Tel:      "admin",
		Birthday: time.Now(),
		Passwd:   passwd,
	}
	return p.Storage([]User{user})
}
