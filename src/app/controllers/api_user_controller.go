package controller

import (
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

	fmt.Println(user)
}
