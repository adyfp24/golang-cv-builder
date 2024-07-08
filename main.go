package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adyfp24/golang-cv-builder/config"
)

func main() {
	config.LoadConfig()
	port := config.GetEnv("SERVER_PORT")
	if port == "" {
		log.Fatal("PORT must be set")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "resume builder with go")
	})

	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
