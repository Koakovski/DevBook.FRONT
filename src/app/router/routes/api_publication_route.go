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
	// LIKE PUBLICATION
	{
		Uri:       "/api/publication/{id}/like",
		Method:    http.MethodPost,
		Handler:   controller.ApiPublicationLikeController,
		IsPrivate: true,
	},
	// UNLIKE PUBLICATION
	{
		Uri:       "/api/publication/{id}/unlike",
		Method:    http.MethodPost,
		Handler:   controller.ApiPublicationUnlikeController,
		IsPrivate: true,
	},
}
