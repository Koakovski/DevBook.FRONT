package controller

import (
	util "devbook-front/src/utils"
	"net/http"
)

func AuthLoginController(w http.ResponseWriter, r *http.Request) {
	util.ExecTemplate(w, "login.html", nil)
}
