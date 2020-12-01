package deletepage

import (
	"net/http"
	"strconv"
	"usermanage/utils"
)

type Deletepage struct {
}

func NewDeletePage() *Deletepage {
	return &Deletepage{}
}

func (this *Deletepage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("Id"))
	var nUsersList utils.UserList
	for _, v := range utils.UsersList {
		if v.Id == id {
			continue
		}
		nUsersList = append(nUsersList, v)
	}
	utils.UsersList = nUsersList
	utils.SaveData()
	http.Redirect(w, r, "/", 302)
}