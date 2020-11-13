package models

import (
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"

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

func (p GobPersistence) fileStore() {
	abs, _ := filepath.Abs(p.Path)

	shuffix := filepath.Ext(abs)
	dir := filepath.Dir(abs)

	search := "*" + shuffix
	searchpath := filepath.Join(dir, search)
	files, _ := filepath.Glob(searchpath)
	fmt.Println(len(files))

	if len(files) > 6 {
		sort.Slice(files, func(i, j int) bool {
			if files[i] < files[j] {
				return true
			}
			return false
		})

		os.Remove(files[1])

	}
	timeStmpstr := strconv.FormatInt(time.Now().Unix(), 10)
	os.Rename(p.Path, strings.Split(p.Path, ".")[0]+timeStmpstr+"."+strings.Split(p.Path, ".")[1])
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

func (c CsvPersistence) fileStore() {
	abs, _ := filepath.Abs(c.Path)

	shuffix := filepath.Ext(abs)
	dir := filepath.Dir(abs)

	search := "*" + shuffix
	searchpath := filepath.Join(dir, search)
	files, _ := filepath.Glob(searchpath)
	fmt.Println(len(files))

	if len(files) > 6 {
		sort.Slice(files, func(i, j int) bool {
			if files[i] < files[j] {
				return true
			}
			return false
		})

		os.Remove(files[1])

	}
	timeStmpstr := strconv.FormatInt(time.Now().Unix(), 10)
	os.Rename(c.Path, strings.Split(c.Path, ".")[0]+timeStmpstr+"."+strings.Split(c.Path, ".")[1])
}

func (c CsvPersistence) LoadUsers() error {
	tmpusers := make([]User, 0, 10)
	file, err := os.Open(c.Path)
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

func (c CsvPersistence) Storage(users []User) error {
	c.fileStore()
	file, err := os.Create(c.Path)
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

func (csv CsvPersistence) InitFile() error {
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
	return csv.Storage([]User{user})
}

type JsonPersistence struct {
	Path string
}

func (js JsonPersistence) LoadUsers() error {
	bytes, err := ioutil.ReadFile(js.Path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &users)
	if err != nil {
		return err
	}
	return nil
}

func (js JsonPersistence) fileStore() {
	abs, _ := filepath.Abs(js.Path)

	shuffix := filepath.Ext(abs)
	dir := filepath.Dir(abs)

	search := "*" + shuffix
	searchpath := filepath.Join(dir, search)
	files, _ := filepath.Glob(searchpath)
	fmt.Println(len(files))

	if len(files) > 6 {
		sort.Slice(files, func(i, j int) bool {
			if files[i] < files[j] {
				return true
			}
			return false
		})

		os.Remove(files[1])

	}
	timeStmpstr := strconv.FormatInt(time.Now().Unix(), 10)
	os.Rename(js.Path, strings.Split(js.Path, ".")[0]+timeStmpstr+"."+strings.Split(js.Path, ".")[1])
}

func (js JsonPersistence) Storage(users []User) error {
	js.fileStore()
	file, err := os.Create(js.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonbytes, err := json.Marshal(users)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonbytes)
	if err != nil {
		return err
	}
	return nil
}

func (js JsonPersistence) InitFile() error {
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
	return js.Storage([]User{user})
}
