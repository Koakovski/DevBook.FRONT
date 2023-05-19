package route

import (
	controller "devbook-front/src/app/controllers"
	"net/http"
)

var HomeRoutes = []Route{
	// HOME PAGE
	{
		Uri:       "/home",
		Method:    http.MethodGet,
		Handler:   controller.HomePageController,
		IsPrivate: true,
	},
}
