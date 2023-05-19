package controller

import (
	util "devbook-front/src/utils"
	"net/http"
)

func UserCreatePageController(w http.ResponseWriter, r *http.Request) {
	util.ExecTemplate(w, "createUser.html", nil)
}

func HomePageController(w http.ResponseWriter, r *http.Request) {
	util.ExecTemplate(w, "home.html", nil)
}

func AuthLoginPageController(w http.ResponseWriter, r *http.Request) {
	util.ExecTemplate(w, "login.html", nil)
}

