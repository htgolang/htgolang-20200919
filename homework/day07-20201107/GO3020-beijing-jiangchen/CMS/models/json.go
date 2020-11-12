package models

import (
	"CMS/consts"
	"os"

	"github.com/adampresley/simdb"
)

//JSONDB ...
//JSONDB struct
type JSONDB struct {
}

//ID ...
func (u User) ID() (string, interface{}) {
	return "name", u.Name
}

//Check ...
// JSONDB Check
func (jsondb *JSONDB) Check() (err error) {
	return nil
}

//Init ...
// JSONDB Init
func (jsondb *JSONDB) Init() (err error) {
	// check if file exists
	// first, check if dir exists
	infodir, err := os.Stat(consts.DBDir + "/" + consts.DBType)
	if os.IsNotExist(err) || !infodir.IsDir() {
		os.MkdirAll(consts.DBDir+"/"+consts.DBType, 0755)
	}
	// second, check if db exists
	infodb, err := os.Stat(consts.DBDir + "/" + consts.DBType + "/User")
	if os.IsNotExist(err) || infodb.IsDir() {
		driver, err := simdb.New(OsFs{}, consts.DBDir+"/"+consts.DBType)
		if err != nil {
			return err
		}
		for _, user := range Users {
			if err = driver.Insert(user); err != nil {
				return err
			}
		}
	}
	return nil
}

//SyncFromDBToMemory ...
// Sync Data from JSONDB to Memory
func (jsondb *JSONDB) SyncFromDBToMemory() (err error) {
	driver, err := simdb.New(OsFs{}, consts.DBDir+"/"+consts.DBType)
	if err != nil {
		return err
	}

	Users = nil

	if _, err = MUE.GetElement(1001); err != nil {
		if err = driver.Open(User{}).Get().AsEntity(MUE.Users); err != nil {
			return err
		}
	} else {
		if err = driver.Open(User{}).Where("id", "!=", 1001).Get().AsEntity(MUE.Users); err != nil {
			return err
		}
	}

	return nil
}

//SyncFromMemoryToDB ...
// Sync Data from Memory to JSONDB
func (jsondb *JSONDB) SyncFromMemoryToDB() (err error) {
	return nil
}

//InsertToDB ...
// Insert Data into JSONDB directly
func (jsondb *JSONDB) InsertToDB(element *User) (err error) {
	driver, err := simdb.New(OsFs{}, consts.DBDir+"/"+consts.DBType)
	if err != nil {
		return err
	}

	if err = driver.Insert(*element); err != nil {
		return err
	}

	return nil
}

//DeleteFromDB ...
// Delete Data from JSONDB directly
func (jsondb *JSONDB) DeleteFromDB(id int) (err error) {
	driver, err := simdb.New(OsFs{}, consts.DBDir+"/"+consts.DBType)
	if err != nil {
		return err
	}

	toDelete := new(User)

	if err = driver.Open(User{}).Where("id", "=", id).First().AsEntity(toDelete); err != nil {
		return err
	}

	if err = driver.Delete(*toDelete); err != nil {
		return err
	}

	return nil
}

//ModifyFromDB ...
// Modify Data to JSONDB directly
func (jsondb *JSONDB) ModifyFromDB(element *User) (err error) {
	driver, err := simdb.New(OsFs{}, consts.DBDir+"/"+consts.DBType)
	if err != nil {
		return err
	}

	id := element.UserID
	toModify := new(User)

	if err = driver.Open(User{}).Where("id", "=", id).First().AsEntity(toModify); err != nil {
		return err
	}

	if err = driver.Delete(*toModify); err != nil {
		return err
	}

	if err = driver.Insert(*element); err != nil {
		return err
	}

	return nil
}

//GetNonAdmin ...
// Get Non-Admin data from JSONDB
func (jsondb *JSONDB) GetNonAdmin() (err error) {
	driver, err := simdb.New(OsFs{}, consts.DBDir+"/"+consts.DBType)
	if err != nil {
		return err
	}

	var us []User
	if err = driver.Open(User{}).Where("name", "!=", "admin").AsEntity(&us); err != nil {
		return err
	}

	*MUE.Users = append(*MUE.Users, us...)

	return nil
}
