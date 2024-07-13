package handlers

import (
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
		Fullname:               r.Form.Get("fullname"),
		Portofolio:             r.Form.Get("portofolio"),
		Email:                  r.Form.Get("email"),
		Phone:                  r.Form.Get("phone"),
		Linkedin:               r.Form.Get("linkedin"),
		Objective:              r.Form.Get("objective"),
		WorkExperience:         r.Form.Get("work-experience"),
		OrganizationExperience: r.Form.Get("organization-experience"),
		Education:              r.Form.Get("education"),
		Certification:          r.Form.Get("certification"),
		Achievement:            r.Form.Get("achievement"),
		City:                   r.Form.Get("city"),
		Skill:                  r.Form.Get("skill"),
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
