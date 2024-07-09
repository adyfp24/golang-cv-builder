package routes

import (
	"net/http"

	"github.com/adyfp24/golang-cv-builder/app/handlers"
)

func InitRoute(){
	http.HandleFunc("/home", handlers.HomeRender)
	http.HandleFunc("/create", handlers.CreateRender)
}