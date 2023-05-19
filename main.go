package main

import (
	"devbook-front/src/app/router"
	"devbook-front/src/infra/config"
	util "devbook-front/src/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	util.LoadTemplates()
	router := router.GenerateRouter()

	stringApiPort := fmt.Sprintf(":%d", config.ApiPort)

	log.Println("Server Listening on", stringApiPort)
	log.Fatal(http.ListenAndServe(stringApiPort, router))
}
