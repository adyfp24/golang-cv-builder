package routes

import (
	"net/http"

	"github.com/adyfp24/golang-cv-builder/app/handlers"
)

func InitRoute(){
	http.HandleFunc("/", handlers.HomeRender)
	http.HandleFunc("/create", handlers.CreateRender)
	http.HandleFunc("/example", handlers.ExampleRender)
	http.HandleFunc("/generate", handlers.GenerateCV)
}