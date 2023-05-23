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
	// FOLLOW USER
	{
		Uri:       "/api/user/{id}/follow",
		Method:    http.MethodPost,
		Handler:   controller.ApiUserFollowController,
		IsPrivate: true,
	},
	// UNFOLLOW USER
	{
		Uri:       "/api/user/{id}/unfollow",
		Method:    http.MethodPost,
		Handler:   controller.ApiUserUnfollowController,
		IsPrivate: true,
	},
	// UPDATE USER
	{
		Uri:       "/api/updateUser",
		Method:    http.MethodPut,
		Handler:   controller.ApiUserUpdateController,
		IsPrivate: true,
	},
	// UPDATE USER PASSWORD
	{
		Uri:       "/api/updatePassword",
		Method:    http.MethodPost,
		Handler:   controller.ApiPasswordUpdateController,
		IsPrivate: true,
	},
}
