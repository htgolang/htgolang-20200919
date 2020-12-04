package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
	"zhao/utils"
)

type Files interface {
	Storage(users []*User) error
	Load() ([]*User, error)
	InitFile() error
}

type phone []byte

func (p phone) String() string {
	return string(p)
}

func ParsePhone(text string) (phone, error) {
	if len(text) == 0 {
		return nil, fmt.Errorf("empty text")
	}
	p := phone(text)
	if len(p) != 11 {
		return nil, fmt.Errorf("11 len numbers")
	}
	for _, v := range p {
		if v < '0' || v > '9' {
			return nil, fmt.Errorf("must is numbers")
		}
	}
	return p, nil
}

type User struct {
	ID       int
	Name     string
	Addr     string
	Tel      string
	Brithday string
	Passwd   string
}

//JSONPersistence json持久化
type JSONPersistence struct {
	name string
	path string
}

func NewJSONPersistence(filename string) *JSONPersistence {
	return &JSONPersistence{
		name: filename,
		path: "../data/",
	}
}

func (j *JSONPersistence) fileStore() error {
	pathAbs, err := filepath.Abs(j.path)
	if err != nil {
		return fmt.Errorf("==[filepath.Abs]== %v", err)
	}
	ext := filepath.Ext(j.name)
	searchfile := "*" + ext
	search := filepath.Join(pathAbs, searchfile)

	matches, err := filepath.Glob(search)
	if err != nil {
		return fmt.Errorf("==[filepath.Golb]== %v", err)
	}

	for len(matches) > 6 {
		sort.Slice(matches, func(i, j int) bool {
			return matches[i] < matches[j]
		})

		err := os.Remove(matches[1])
		if err != nil {
			return fmt.Errorf("==[os.Remove]== %v", err)
		}
	}
	timeStmpStr := strconv.FormatInt(time.Now().Unix(), 10)

	os.Rename(filepath.Join(pathAbs, j.name), filepath.Join(pathAbs, strings.Split(j.name, ".")[0]+timeStmpStr+ext))
	return nil
}

func (j *JSONPersistence) Storage(users []*User) error {
	err := j.fileStore()
	if err != nil {
		return err
	}

	jsbytes, err := json.Marshal(users)
	if err != nil {
		return fmt.Errorf("==[json.Marshal]== %v", err)
	}

	file, err := os.Create("../data/" + j.name)
	if err != nil {
		return fmt.Errorf("==[os.Create]== %v", err)
	}
	defer file.Close()

	_, err = file.Write(jsbytes)
	if err != nil {
		return fmt.Errorf("==[file.Write]== %v", err)
	}
	return nil
}

func (j *JSONPersistence) Load() ([]*User, error) {
	jsbytes, err := ioutil.ReadFile("../data/" + j.name)
	if err != nil {
		return nil, fmt.Errorf("==[Load--os.Open]== %v", err)
	}

	users := new([]*User)
	err = json.Unmarshal(jsbytes, users)
	if err != nil {
		return nil, fmt.Errorf("==[json.Unmarshal]== %v", err)
	}

	return *users, nil
}

func (j *JSONPersistence) InitFile() error {
	passwd := utils.Md5Convert("admin")
	user := &User{
		ID:       0,
		Name:     "admin",
		Addr:     "admin",
		Tel:      "00000000000",
		Brithday: time.Now().Format("2006/01/02"),
		Passwd:   passwd,
	}

	return j.Storage([]*User{user})
}
