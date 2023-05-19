package controller

import (
	"devbook-front/src/infra/config"
	request "devbook-front/src/infra/requests"
	util "devbook-front/src/utils"
	"fmt"
	"net/http"
)

func UserCreatePageController(w http.ResponseWriter, r *http.Request) {
	util.ExecTemplate(w, "createUser.html", nil)
}

func AuthLoginPageController(w http.ResponseWriter, r *http.Request) {
	util.ExecTemplate(w, "login.html", nil)
}

func HomePageController(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publication", config.ApiUrl)
	response, err := request.RequestWithAuth(r, http.MethodGet, url, nil)

	fmt.Println(response, err)

	util.ExecTemplate(w, "home.html", response)
}
