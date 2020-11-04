# **CMS Version 3**

基于原项目[CMS Version 2](https://github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/GO3020-beijing-jiangchen), 将用户信息持久化到boltdb数据库中, 有关boltdb数据库详情请参见[github项目](https://github.com/boltdb/bolt)

* ## **新增boltdb相关操作的函数如下所示**

```go
package models

import (
	"CMS/consts"

	"github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"
)

//SyncBoltDBToMem ...
// Sync Data from BoltDB to Memory
func SyncBoltDBToMem() (err error) {
	store, err := bolthold.Open(consts.DBPath, 0666, nil)
	defer store.Close()

	if err != nil {
		return err
	}

	// empty Users slice before Sync
	Users = nil

	if err = store.Find(&Users, bolthold.Where("ID").Ge(0)); err != nil {
		return err
	}

	return nil
}

//SyncMemToBoltDB ...
// Sync Data from Memory to BoltDB
func SyncMemToBoltDB() (err error) {
	store, err := bolthold.Open(consts.DBPath, 0666, nil)
	defer store.Close()

	if err != nil {
		return err
	}

	// insert all data in one transaction
	err = store.Bolt().Update(func(tx *bolt.Tx) error {
		for i := range Users {
			_ = store.TxInsert(tx, Users[i].ID, Users[i])
		}
		return nil
	})

	return nil
}

//InsertToBoltDB ...
// Insert Data into BoltDB directly
func InsertToBoltDB(element *User) (err error) {
	store, err := bolthold.Open(consts.DBPath, 0666, nil)
	defer store.Close()

	if err != nil {
		return err
	}

	err = store.Insert(element.ID, element)
	if err != nil {
		return err
	}

	return nil
}

//DeleteFromBoltDB ...
// Delete Data from BoltDB directly
func DeleteFromBoltDB(id int) (err error) {
	store, err := bolthold.Open(consts.DBPath, 0666, nil)
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

//ModifyFromBoltDB ...
// Modify Data to BoltDB directly
func ModifyFromBoltDB(element *User) (err error) {
	store, err := bolthold.Open(consts.DBPath, 0666, nil)
	defer store.Close()

	if err != nil {
		return err
	}

	err = store.Update(element.ID, element)
	if err != nil {
		return err
	}

	return nil
}

//GetNonAdmin ...
// Get Non-Admin data from BoltDB
func GetNonAdmin() (err error) {
	store, err := bolthold.Open(consts.DBPath, 0666, nil)
	defer store.Close()

	if err != nil {
		return err
	}

	// Append all non-admin user to Users slice
	if err = store.Find(&Users, bolthold.Where("ID").Ge(0).And("Name").Ne("admin")); err != nil {
		return err
	}

	return nil
}
```
