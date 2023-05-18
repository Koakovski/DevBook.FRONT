package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running WebApp...")

	log.Fatal(http.ListenAndServe(":5000", nil))
}
