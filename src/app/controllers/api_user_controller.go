package controller

import (
	"bytes"
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
		log.Fatal(err)
	}

	response, err := http.Post("http://localhost:3000/user", "application/json", bytes.NewBuffer(user))
	if err != nil {
		log.Fatal(err)
	}
	response.Body.Close()

	fmt.Println(response.Body)
}
