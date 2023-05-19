package middlware

import (
	cookie "devbook-front/src/infra/cookies"
	"net/http"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := cookie.ReadCookie(r); err != nil {
			http.Redirect(w, r, "login", http.StatusMovedPermanently)
			return
		}

		next(w, r)
	}
}
