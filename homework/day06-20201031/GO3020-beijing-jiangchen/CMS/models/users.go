package models

import (
	"log"
	"time"

	"github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"

	"CMS/consts"
)

//User ...
//Global data structure of User element.
type User struct {
	ID       int
	Name     string `boltholdIndex:"Name"`
	Tel      string
	Address  string
	Birthday time.Time
	Password [16]byte
}

//Users ...
//Global Data Structure, for each element of Users slice,
//in other words, User, has the same data structure below:
/*
	ID int
	Name string
	Tel string
	Address string
	Birthday time.Time
	password [16]byte
*/
var Users []User

//LoginCount ...
// Login count number
var LoginCount int = 0

func init() {
	Users = make([]User, 0)
	Users = append(Users, *GenerateElement(1001, "admin", "+1 4406665321", "2426 Wildwood Street, Medina, Ohio", "1990-01-01", "admin"))
	boltInit()
}

func boltInit() {
	store, err := bolthold.Open(consts.DBPath, 0666, nil)
	defer store.Close()

	if err != nil {
		// handle error
		log.Fatal(err)
	}

	// insert all data in one transaction
	err = store.Bolt().Update(func(tx *bolt.Tx) error {
		for i := range Users {
			err := store.TxInsert(tx, Users[i].ID, Users[i])
			if err != nil {
				return err
			}
		}
		return nil
	})
}
