package controller

import (
	"bytes"
	presenter "devbook-front/src/app/presenters"
	"devbook-front/src/infra/config"
	request "devbook-front/src/infra/requests"
	"encoding/json"
	"fmt"
	"net/http"
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
