package funcs

import (
	"CMS/models"
	"fmt"
)

//SyncToMem ...
// Wrapper of SyncBoltDBToMem.
func SyncToMem() {
	err := models.SyncBoltDBToMem()
	if err != nil {
		fmt.Println("SyncMem Failed...")
		fmt.Printf("Error Message: %v\n", err)
	}
}

//SyncToDB ...
// Wrapper of SyncMemToBoltDB.
func SyncToDB() {
	err := models.SyncMemToBoltDB()
	if err != nil {
		fmt.Println("SyncDB Failed...")
		fmt.Printf("Error Message: %v\n", err)
	}
}
