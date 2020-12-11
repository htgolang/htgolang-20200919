package controllers

import (
	"net/http"
	"strconv"
	"usermanage/services"
)

func DeletePage(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("Id"))
	services.DeleteUser(id)
	http.Redirect(w, r, "/", 302)
}
