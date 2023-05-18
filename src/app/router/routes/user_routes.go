package route

import (
	controller "devbook-front/src/app/controllers"
	"net/http"
)

var UserRoutes = []Route{
	// LOAD CREATE USER PAGE
	{
		Uri:       "/createUser",
		Method:    http.MethodGet,
		Handler:   controller.UserCreatePageController,
		IsPrivate: false,
	},
}
