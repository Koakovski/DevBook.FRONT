package router

import (
	route "devbook-front/src/app/router/routes"

	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	return route.ConfigureRoutes(mux.NewRouter())
}
