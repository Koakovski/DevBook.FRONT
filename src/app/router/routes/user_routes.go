package route

import (
	controller "devbook-front/src/app/controllers"
	"net/http"
)

var UserRoutes = []Route{
	// LOGIN
	{
		Uri:       "/createUser",
		Method:    http.MethodGet,
		Handler:   controller.UserCreatePageController,
		IsPrivate: false,
	},
}
