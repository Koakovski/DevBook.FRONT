package route

import (
	controller "devbook-front/src/app/controllers"
	"net/http"
)

var PagesRoutes = []Route{
	// LOAD LOGIN PAGE
	{
		Uri:       "/",
		Method:    http.MethodGet,
		Handler:   controller.AuthLoginPageController,
		IsPrivate: false,
	},
	{
		Uri:       "/login",
		Method:    http.MethodGet,
		Handler:   controller.AuthLoginPageController,
		IsPrivate: false,
	},
	// LOAD HOME PAGE
	{
		Uri:       "/home",
		Method:    http.MethodGet,
		Handler:   controller.HomePageController,
		IsPrivate: true,
	},
	// LOAD CREATE USER PAGE
	{
		Uri:       "/createUser",
		Method:    http.MethodGet,
		Handler:   controller.UserCreatePageController,
		IsPrivate: false,
	},
	// LOAD UPDATE PUBLICATION PAGE
	{
		Uri:       "/publication/{id}/update",
		Method:    http.MethodGet,
		Handler:   controller.PublicationUpdatePageController,
		IsPrivate: false,
	},
}
