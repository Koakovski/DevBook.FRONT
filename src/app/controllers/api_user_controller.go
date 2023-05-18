package controller

import (
	"bytes"
	presenter "devbook-front/src/app/presenters"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ApiUserCreateController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nickName": r.FormValue("nickname"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	response, err := http.Post("http://localhost:3000/user", "application/json", bytes.NewBuffer(user))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	presenter.ReponsePresenter(w, response.StatusCode, nil)
}
