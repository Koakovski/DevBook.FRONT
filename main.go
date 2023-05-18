package main

import (
	"devbook-front/src/app/router"
	"log"
	"net/http"
)

func main() {
	router := router.GenerateRouter()

	stringApiPort := ":5000"

	log.Println("Server Listening on", stringApiPort)
	log.Fatal(http.ListenAndServe(stringApiPort, router))
}
