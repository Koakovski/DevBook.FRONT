package route

import (
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
	routes := append(AuthRoutes, UserRoutes...)

	for _, route := range routes {
		router.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
