package route

import (
	controller "devbook-front/src/app/controllers"
	"net/http"
)

var AuthRoutes = []Route{
	// LOGIN
	{
		Uri:       "/",
		Method:    http.MethodGet,
		Handler:   controller.AuthLoginController,
		IsPrivate: false,
	},
	{
		Uri:       "/login",
		Method:    http.MethodGet,
		Handler:   controller.AuthLoginController,
		IsPrivate: false,
	},
}
