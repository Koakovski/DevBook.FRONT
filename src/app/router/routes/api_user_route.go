package route

import (
	controller "devbook-front/src/app/controllers"
	"net/http"
)

var ApiUserRoutes = []Route{
	// REQUEST CREATE USER
	{
		Uri:       "/api/user",
		Method:    http.MethodPost,
		Handler:   controller.ApiUserCreateController,
		IsPrivate: false,
	},
	// AUTH LOGIN
	{
		Uri:       "/api/login",
		Method:    http.MethodPost,
		Handler:   controller.ApiAuthLoginController,
		IsPrivate: false,
	},
}
