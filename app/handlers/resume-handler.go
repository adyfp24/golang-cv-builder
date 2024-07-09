package handlers

import (
	"html/template"
	"net/http"

	"github.com/adyfp24/golang-cv-builder/app/models"
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

func GenerateCV(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Ambil data dari form
    r.ParseForm()
    cv := models.Resume{
        Name:       r.Form.Get("name"),
        Email:      r.Form.Get("email"),
        Phone:      r.Form.Get("phone"),
        Experience: r.Form.Get("experience"),
        Education:  r.Form.Get("education"),
    }

    // Generate PDF CV
    pdfBytes, err := service.GeneratePDF(cv)
    if err != nil {
        http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
        return
    }

    // Kirim PDF sebagai response
    w.Header().Set("Content-Type", "application/pdf")
    w.Header().Set("Content-Disposition", "attachment; filename=cv.pdf")
    w.Write(pdfBytes)
}
