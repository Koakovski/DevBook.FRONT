package presenter

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	Error string `json:"error"`
}

func ErrorPresenter(w http.ResponseWriter, r *http.Response) {
	var apiErr ApiError
	json.NewDecoder(r.Body).Decode(&apiErr)

	ReponsePresenter(w, r.StatusCode, apiErr)
}
