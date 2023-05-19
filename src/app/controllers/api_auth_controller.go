package controller

import (
	"bytes"
	presenter "devbook-front/src/app/presenters"
	model "devbook-front/src/domain/models"
	"devbook-front/src/infra/config"
	cookie "devbook-front/src/infra/cookies"
	"encoding/json"
	"fmt"
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

	url := fmt.Sprintf("%s/login", config.ApiUrl)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(loginData))
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

	if err := cookie.SaveCookie(w, authData.UserId, authData.Token); err != nil {
		presenter.ReponsePresenter(w, http.StatusInternalServerError, presenter.ApiError{Error: err.Error()})
		return
	}

	presenter.ReponsePresenter(w, http.StatusOK, nil)
}
