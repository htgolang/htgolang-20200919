package funcs

import (
	"CMS/models"
	"fmt"
)

//SyncToMem ...
// Wrapper of SyncBoltDBToMem.
func SyncToMem() {
	err := models.MUE.PersistentStorage.SyncFromDBToMemory()
	if err != nil {
		fmt.Println("SyncMem Failed...")
		fmt.Printf("Error Message: %v\n", err)
	}
}

//SyncToDB ...
// Wrapper of SyncMemToBoltDB.
func SyncToDB() {
	err := models.MUE.PersistentStorage.SyncFromMemoryToDB()
	if err != nil {
		fmt.Println("SyncDB Failed...")
		fmt.Printf("Error Message: %v\n", err)
	}
}
