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

func ApiPublicationCreateController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publicationData, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if err != nil {
		presenter.ReponsePresenter(w, http.StatusBadRequest, presenter.ApiError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publication", config.ApiUrl)
	response, err := request.RequestWithAuth(r, http.MethodPost, url, bytes.NewBuffer(publicationData))
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

func ApiPublicationLikeController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		presenter.ReponsePresenter(w, http.StatusBadRequest, presenter.ApiError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publication/%d/like", config.ApiUrl, publicationId)
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
