package controller

import "net/http"

func AuthLoginController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tela de Login"))
}
