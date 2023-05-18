package main

import (
	"devbook-front/src/app/router"
	util "devbook-front/src/utils"
	"log"
	"net/http"
)

func main() {
	util.LoadTemplates()
	router := router.GenerateRouter()

	stringApiPort := ":5000"

	log.Println("Server Listening on", stringApiPort)
	log.Fatal(http.ListenAndServe(stringApiPort, router))
}
