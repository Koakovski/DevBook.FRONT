package controller

import (
	"bytes"
	presenter "devbook-front/src/app/presenters"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ApiAuthLoginController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	loginData, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		presenter.ReponsePresenter(w, http.StatusBadRequest, presenter.ApiError{Error: err.Error()})
		return
	}

	response, err := http.Post("http://localhost:3000/login", "application/json", bytes.NewBuffer(loginData))
	if err != nil {
		presenter.ReponsePresenter(w, http.StatusInternalServerError, presenter.ApiError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		presenter.ErrorPresenter(w, response)
		return
	}

	token, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(token)
}
