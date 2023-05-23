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
	// UPDATE PUBLICATION
	{
		Uri:       "/api/publication/{id}",
		Method:    http.MethodPut,
		Handler:   controller.ApiPublicationUpdateController,
		IsPrivate: true,
	},
	// DELETE PUBLICATION
	{
		Uri:       "/api/publication/{id}",
		Method:    http.MethodDelete,
		Handler:   controller.ApiPublicationDeleteController,
		IsPrivate: true,
	},
}
