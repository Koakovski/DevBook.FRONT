package main

import (
	"devbook-front/src/app/router"
	"devbook-front/src/infra/config"
	cookie "devbook-front/src/infra/cookies"
	util "devbook-front/src/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()
	cookie.ConfigureCookie()

	util.LoadTemplates()
	router := router.GenerateRouter()

	stringApiPort := fmt.Sprintf(":%d", config.ApiPort)

	log.Println("Server Listening on", stringApiPort)
	log.Fatal(http.ListenAndServe(stringApiPort, router))
}
