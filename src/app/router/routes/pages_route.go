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
	// LOAD CREATE USER PAGE
	{
		Uri:       "/createUser",
		Method:    http.MethodGet,
		Handler:   controller.UserCreatePageController,
		IsPrivate: false,
	},
	// LOAD HOME PAGE
	{
		Uri:       "/home",
		Method:    http.MethodGet,
		Handler:   controller.HomePageController,
		IsPrivate: true,
	},
	// LOAD UPDATE PUBLICATION PAGE
	{
		Uri:       "/publication/{id}/update",
		Method:    http.MethodGet,
		Handler:   controller.PublicationUpdatePageController,
		IsPrivate: true,
	},
	// LOAD SEARCH USERS PAGE
	{
		Uri:       "/searchUsers",
		Method:    http.MethodGet,
		Handler:   controller.SearchUsersPageController,
		IsPrivate: true,
	},
	// LOAD USER PROFILE PAGE
	{
		Uri:       "/user/{id}",
		Method:    http.MethodGet,
		Handler:   controller.UserProfilePageController,
		IsPrivate: true,
	},
	// LOAD USER PROFILE PAGE
	{
		Uri:       "/profile",
		Method:    http.MethodGet,
		Handler:   controller.UserAuthenticatedProfilePageController,
		IsPrivate: true,
	},
}
