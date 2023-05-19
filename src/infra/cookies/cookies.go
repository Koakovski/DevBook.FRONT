package cookie

import (
	"devbook-front/src/infra/config"
	"net/http"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func ConfigureCookie() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func SaveCookie(w http.ResponseWriter, Id, token string) error {
	data := map[string]string{
		"id":    Id,
		"token": token,
	}

	cookieIdentifier := "AuthData"

	codedData, err := s.Encode(cookieIdentifier, data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     cookieIdentifier,
		Value:    codedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}
