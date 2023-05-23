package controller

import (
	"bytes"
	presenter "devbook-front/src/app/presenters"
	"devbook-front/src/infra/config"
	request "devbook-front/src/infra/requests"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func ApiUserFollowController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		presenter.ReponsePresenter(w, http.StatusBadRequest, presenter.ApiError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/user/%d/follow", config.ApiUrl, userId)
	response, err := request.RequestWithAuth(r, http.MethodPost, url, nil)
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

func ApiUserUnfollowController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		presenter.ReponsePresenter(w, http.StatusBadRequest, presenter.ApiError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/user/%d/unfollow", config.ApiUrl, userId)
	response, err := request.RequestWithAuth(r, http.MethodPost, url, nil)
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
