package handlers

import (
	"html/template"
	"net/http"
)

func HomeRender(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

	tmpl := template.Must(template.ParseFiles("app/views/home.html"))
	tmpl.Execute(w, nil)
}

func CreateRender(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

	tmpl := template.Must(template.ParseFiles("app/views/create.html"))
	tmpl.Execute(w, nil)
}

