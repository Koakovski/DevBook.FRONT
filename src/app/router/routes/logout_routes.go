package route

import (
	controller "devbook-front/src/app/controllers"
	"net/http"
)

var LogoutRoutes = []Route{
	// LOGOUT
	{
		Uri:       "/logout",
		Method:    http.MethodGet,
		Handler:   controller.LogoutController,
		IsPrivate: true,
	},
}
