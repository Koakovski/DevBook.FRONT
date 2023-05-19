package controller

import (
	"bytes"
	presenter "devbook-front/src/app/presenters"
	model "devbook-front/src/domain/models"
	"encoding/json"
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

	var authData model.AuthData
	if err := json.NewDecoder(response.Body).Decode(&authData); err != nil {
		presenter.ReponsePresenter(w, http.StatusUnprocessableEntity, presenter.ApiError{Error: err.Error()})
		return
	}

	presenter.ReponsePresenter(w, http.StatusOK, nil)
}
