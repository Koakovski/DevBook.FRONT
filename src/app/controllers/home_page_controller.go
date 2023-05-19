package controller

import (
	util "devbook-front/src/utils"
	"net/http"
)

func HomePageController(w http.ResponseWriter, r *http.Request) {
	util.ExecTemplate(w, "home.html", nil)
}
