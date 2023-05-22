package route

import (
	controller "devbook-front/src/app/controllers"
	"net/http"
)

var ApiPublicationRoutes = []Route{
	// CREATE PUBLICATION
	{
		Uri:       "/api/publication",
		Method:    http.MethodPost,
		Handler:   controller.ApiPublicationCreateController,
		IsPrivate: true,
	},
}
