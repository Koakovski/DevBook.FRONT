package cookie

import (
	"devbook-front/src/infra/config"
	"net/http"

	"github.com/gorilla/securecookie"
)

var CookieIdentifier = "AuthData"
var s *securecookie.SecureCookie

func ConfigureCookie() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func SaveCookie(w http.ResponseWriter, Id, token string) error {
	data := map[string]string{
		"id":    Id,
		"token": token,
	}

	codedData, err := s.Encode(CookieIdentifier, data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     CookieIdentifier,
		Value:    codedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

func ReadCookie(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie(CookieIdentifier)
	if err != nil {
		return nil, err
	}

	cookieValues := make(map[string]string)

	if err := s.Decode(CookieIdentifier, cookie.Value, &cookieValues); err != nil {
		return nil, err
	}

	return cookieValues, nil
}
