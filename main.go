package main

import (
	"log"
	"net/http"

	"github.com/adyfp24/golang-cv-builder/app/routes"
	"github.com/adyfp24/golang-cv-builder/config"
)

func main() {
	routes.InitRoute()
	config.LoadConfig()
	port := config.GetEnv("SERVER_PORT")

	if port == "" {
		log.Fatal("PORT must be set")
	}

	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
