package services

import (
	"userMgr/models"
)

func Admin() bool {
	user, _ := IDFindUser(1)
	var role = models.Role{
		User:      user,
		RolePerm:  31,
		RoleTitle: "",
	}
	if IsAdmin(&role) {
		return true
	}
	return false
}

func IsAdmin(r *models.Role) bool {
	if r.RolePerm == models.PermAdmin {
		r.RoleTitle = "Admin"
		return true
	}
	return false
}

func IsSuperUser(r *models.Role) bool {
	if r.RolePerm == models.PermSuper {
		r.RoleTitle = "SuperUser"
		return true
	}
	return false
}

func IsViewer(r *models.Role) bool {
	if r.RolePerm == models.PermUser {
		r.RoleTitle = "User"
		return true
	}
	return false
}
