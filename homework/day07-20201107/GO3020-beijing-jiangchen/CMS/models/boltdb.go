package models

import (
	"CMS/consts"
	"fmt"
	"log"
	"os"

	"github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"
)

//BoltDB ...
// BoltDB struct
type BoltDB struct {
}

//Check ...
// BoltDB Check
func (boltdb *BoltDB) Check() (err error) {
	return nil
}

//Init ...
// BoltDB initialization
func (boltdb *BoltDB) Init() (err error) {
	// check if file exists
	info, err := os.Stat(consts.DBDir + "/" + consts.DBType + "/Users.db")
	if err != nil {
		fmt.Printf("boltdb %s does not exists, create a new db\n", consts.DBDir+"/"+consts.DBType+"/Users.db")
	}
	os.MkdirAll(consts.DBDir+"/"+consts.DBType, 0755)

	if os.IsNotExist(err) || info.IsDir() {
		store, err := bolthold.Open(consts.DBDir+"/"+consts.DBType+"/Users.db", 0666, nil)
		defer store.Close()
		if err != nil {
			// handle error
			log.Fatal(err)
		}
		// insert all data in one transaction
		err = store.Bolt().Update(func(tx *bolt.Tx) error {
			for i := range Users {
				err := store.TxInsert(tx, (*MUE.Users)[i].UserID, (*MUE.Users)[i])
				if err != nil {
					return err
				}
			}
			return nil
		})
	}
	return nil
}

//SyncFromDBToMemory ...
// Sync Data from BoltDB to Memory
func (boltdb *BoltDB) SyncFromDBToMemory() (err error) {
	store, err := bolthold.Open(consts.DBDir+"/"+consts.DBType+"/Users.db", 0666, nil)
	defer store.Close()

	if err != nil {
		return err
	}

	// empty Users slice before Sync
	Users = nil

	if err = store.Find(MUE.Users, bolthold.Where("UserID").Ge(0)); err != nil {
		return err
	}

	return nil
}

//SyncFromMemoryToDB ...
// Sync Data from Memory to BoltDB
func (boltdb *BoltDB) SyncFromMemoryToDB() (err error) {
	store, err := bolthold.Open(consts.DBDir+"/"+consts.DBType+"/Users.db", 0666, nil)
	defer store.Close()

	if err != nil {
		return err
	}

	// insert all data in one transaction
	err = store.Bolt().Update(func(tx *bolt.Tx) error {
		for i := range *MUE.Users {
			_ = store.TxInsert(tx, (*MUE.Users)[i].UserID, (*MUE.Users)[i])
		}
		return nil
	})

	return nil
}

//InsertToDB ...
// Insert Data into BoltDB directly
func (boltdb *BoltDB) InsertToDB(element *User) (err error) {
	store, err := bolthold.Open(consts.DBDir+"/"+consts.DBType+"/Users.db", 0666, nil)
	defer store.Close()

	if err != nil {
		return err
	}

	err = store.Insert(element.UserID, element)
	if err != nil {
		return err
	}

	return nil
}

//DeleteFromDB ...
// Delete Data from BoltDB directly
func (boltdb *BoltDB) DeleteFromDB(id int) (err error) {
	store, err := bolthold.Open(consts.DBDir+"/"+consts.DBType+"/Users.db", 0666, nil)
	defer store.Close()

	if err != nil {
		return err
	}

	err = store.Delete(id, &User{})
	if err != nil {
		return err
	}

	return nil
}

//ModifyFromDB ...
// Modify Data to BoltDB directly
func (boltdb *BoltDB) ModifyFromDB(element *User) (err error) {
	store, err := bolthold.Open(consts.DBDir+"/"+consts.DBType+"/Users.db", 0666, nil)
	defer store.Close()

	if err != nil {
		return err
	}

	err = store.Update(element.UserID, element)
	if err != nil {
		return err
	}

	return nil
}

//GetNonAdmin ...
// Get Non-Admin data from BoltDB
func (boltdb *BoltDB) GetNonAdmin() (err error) {
	store, err := bolthold.Open(consts.DBDir+"/"+consts.DBType+"/Users.db", 0666, nil)
	defer store.Close()

	if err != nil {
		return err
	}

	// Append all non-admin user to Users slice
	if err = store.Find(MUE.Users, bolthold.Where("UserID").Ge(0).And("Name").Ne("admin")); err != nil {
		return err
	}

	return nil
}
