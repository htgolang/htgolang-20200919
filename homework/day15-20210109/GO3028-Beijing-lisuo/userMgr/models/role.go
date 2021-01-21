package models

const (
	// 0 0 0 0 0
	// v c d m q
	PermAdmin uint = 31
	PermSuper uint = 27
	PermUser  uint = 17
)

type Role struct {
	*User
	RolePerm  uint
	RoleTitle string
}

//type RoleTitle struct {
//	Admin      string
//	SuperUser  string
//	User string
//	Viewer     string
//}

type RolePerm struct {
	View   bool //
	Create bool // UserController.Create
	Delete bool // UserController.Delete
	Modify bool // UserController.Edit
	Query  bool // UserController.Query
}
