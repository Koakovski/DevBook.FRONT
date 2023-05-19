package controller

import (
	"bytes"
	presenter "devbook-front/src/app/presenters"
	"devbook-front/src/infra/config"
	"encoding/json"
	"fmt"
	"net/http"
)

func ApiUserCreateController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	userData, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nickName": r.FormValue("nickName"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		presenter.ReponsePresenter(w, http.StatusBadRequest, presenter.ApiError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/user", config.ApiUrl)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(userData))
	if err != nil {
		presenter.ReponsePresenter(w, http.StatusInternalServerError, presenter.ApiError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		presenter.ErrorPresenter(w, response)
		return
	}

	presenter.ReponsePresenter(w, response.StatusCode, nil)
}
