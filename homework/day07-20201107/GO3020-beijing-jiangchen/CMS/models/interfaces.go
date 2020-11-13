package models

//Storage ...
// Universal Storage interface
type Storage interface {
	Check() error
	Init() error
	SyncFromDBToMemory() error
	SyncFromMemoryToDB() error
	InsertToDB(*User) error
	DeleteFromDB(int) error
	ModifyFromDB(*User) error
	GetNonAdmin() error
}
