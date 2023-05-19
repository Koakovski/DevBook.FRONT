package request

import (
	cookie "devbook-front/src/infra/cookies"
	"fmt"
	"io"
	"net/http"
)

func RequestWithAuth(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	authCookie, _ := cookie.ReadCookie(r)

	bearerToken := fmt.Sprintf("Bearer %s", authCookie["token"])
	request.Header.Add("Authorization", bearerToken)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
