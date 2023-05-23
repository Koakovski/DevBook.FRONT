package controller

import (
	cookie "devbook-front/src/infra/cookies"
	"net/http"
)

func LogoutController(w http.ResponseWriter, r *http.Request) {
	cookie.DeleteCookie(w)

	http.Redirect(w, r, "/login", http.StatusMovedPermanently)
}
