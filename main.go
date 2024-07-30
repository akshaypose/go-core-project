package main

import (
	"go-core-project/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.SetupRoutes()

	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
