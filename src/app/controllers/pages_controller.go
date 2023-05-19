package controller

import (
	presenter "devbook-front/src/app/presenters"
	model "devbook-front/src/domain/models"
	"devbook-front/src/infra/config"
	request "devbook-front/src/infra/requests"
	util "devbook-front/src/utils"
	"encoding/json"
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
	if err != nil {
		presenter.ReponsePresenter(w, http.StatusInternalServerError, presenter.ApiError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		presenter.ErrorPresenter(w, response)
		return
	}

	var publications []model.Publication

	if err := json.NewDecoder(response.Body).Decode(&publications); err != nil {
		presenter.ReponsePresenter(w, http.StatusInternalServerError, presenter.ApiError{Error: err.Error()})
		return
	}

	util.ExecTemplate(w, "home.html", struct {
		Publications []model.Publication
	}{
		Publications: publications,
	})
}
