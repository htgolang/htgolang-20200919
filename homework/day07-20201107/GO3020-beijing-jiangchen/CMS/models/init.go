package models

func init() {
	Users = make([]User, 0)
	Users = append(Users, *GenerateElement(1001, "admin", "+1 4406665321", "2426 Wildwood Street, Medina, Ohio", "1990-01-01", "admin"))
	MUE = new(MemoryUsersEntity)
}
