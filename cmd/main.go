package main

import (
	"github.com/loveleshsharma/rate-limiter/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.SetupRoutes()

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("error starting server: ", err)
	}
}
