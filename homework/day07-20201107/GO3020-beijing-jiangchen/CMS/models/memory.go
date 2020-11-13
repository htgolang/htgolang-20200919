package models

//MemoryDB ...
// actually non...
type MemoryDB struct {
}

//Check ...
//MemDB Check
func (memdb *MemoryDB) Check() (err error) {
	return nil
}

//Init ...
// MemDB initialization
func (memdb *MemoryDB) Init() (err error) {
	return nil
}

//SyncFromDBToMemory ...
// Sync Data from MemDB to Memory
func (memdb *MemoryDB) SyncFromDBToMemory() (err error) {
	return nil
}

//SyncFromMemoryToDB ...
// Sync Data from Memory to MemoryDB
func (memdb *MemoryDB) SyncFromMemoryToDB() (err error) {
	return nil
}

//InsertToDB ...
// Insert Data into MemDB directly
func (memdb *MemoryDB) InsertToDB(element *User) (err error) {
	return nil
}

//DeleteFromDB ...
// Delete Data from MemDB directly
func (memdb *MemoryDB) DeleteFromDB(id int) (err error) {
	return nil
}

//ModifyFromDB ...
// Modify Data to MemDB directly
func (memdb *MemoryDB) ModifyFromDB(element *User) (err error) {
	return nil
}

//GetNonAdmin ...
// Get Non-Admin data from MemDB
func (memdb *MemoryDB) GetNonAdmin() (err error) {
	return nil
}
