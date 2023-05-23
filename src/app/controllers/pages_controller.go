package controller

import (
	presenter "devbook-front/src/app/presenters"
	model "devbook-front/src/domain/models"
	"devbook-front/src/infra/config"
	cookie "devbook-front/src/infra/cookies"
	request "devbook-front/src/infra/requests"
	util "devbook-front/src/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UserCreatePageController(w http.ResponseWriter, r *http.Request) {
	util.ExecTemplate(w, "createUser.html", nil)
}

func AuthLoginPageController(w http.ResponseWriter, r *http.Request) {
	cookieData, _ := cookie.ReadCookie(r)

	if cookieData["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusMovedPermanently)
		return
	}

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

	cookie, _ := cookie.ReadCookie(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	util.ExecTemplate(w, "home.html", struct {
		Publications []model.Publication
		UserId       uint64
	}{
		Publications: publications,
		UserId:       userId,
	})
}

func PublicationUpdatePageController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		presenter.ReponsePresenter(w, http.StatusBadRequest, presenter.ApiError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publication/%d", config.ApiUrl, publicationId)
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

	var publication model.Publication

	if err := json.NewDecoder(response.Body).Decode(&publication); err != nil {
		presenter.ReponsePresenter(w, http.StatusInternalServerError, presenter.ApiError{Error: err.Error()})
		return
	}

	util.ExecTemplate(w, "updatePublication.html", publication)
}
