package route

import (
	middlware "devbook-front/src/app/middlwares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri       string
	Method    string
	Handler   func(http.ResponseWriter, *http.Request)
	IsPrivate bool
}

func ConfigureRoutes(router *mux.Router) *mux.Router {
	routes := PagesRoutes

	routes = append(routes, ApiUserRoutes...)
	routes = append(routes, ApiPublicationRoutes...)
	routes = append(routes, LogoutRoutes...)

	for _, route := range routes {
		if route.IsPrivate {
			router.HandleFunc(route.Uri, middlware.Authenticate(route.Handler)).Methods(route.Method)
		} else {
			router.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
