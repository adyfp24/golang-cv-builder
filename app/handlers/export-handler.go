package handlers

import(
	"github.com/adyfp24/golang-cv-builder/app/models"
	"github.com/adyfp24/golang-cv-builder/app/services"
	"net/http"
)

func GenerateCV(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    r.ParseForm()
    cv := models.Resume{
        Name:       r.Form.Get("name"),
        Email:      r.Form.Get("email"),
        Phone:      r.Form.Get("phone"),
        Experience: r.Form.Get("experience"),
        Education:  r.Form.Get("education"),
    }

    pdfBytes, err := services.GeneratePDF(cv)
    if err != nil {
        http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/pdf")
    w.Header().Set("Content-Disposition", "attachment; filename=cv.pdf")
    w.Write(pdfBytes)
}
