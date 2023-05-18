package controller

import (
	util "devbook-front/src/utils"
	"net/http"
)

func UserCreatePageController(w http.ResponseWriter, r *http.Request) {
	util.ExecTemplate(w, "createUser.html", nil)
}
