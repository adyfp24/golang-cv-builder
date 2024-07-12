package services

import (
	"bytes"

	"github.com/adyfp24/golang-cv-builder/app/models"
	"github.com/jung-kurt/gofpdf"
)

func GeneratePDF(cv models.Resume) ([]byte, error) {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()

    pdf.SetFont("Arial", "B", 16)
    pdf.CellFormat(0, 10, "Curriculum Vitae", "", 1, "C", false, 0, "")

    pdf.SetFont("Arial", "", 12)
    pdf.CellFormat(0, 10, "Personal Information", "", 1, "L", false, 0, "")
    pdf.Ln(5)
    pdf.CellFormat(0, 10, "Name: "+cv.FirstName+" "+cv.LastName, "", 1, "L", false, 0, "")
    pdf.CellFormat(0, 10, "Email: "+cv.Email, "", 1, "L", false, 0, "")
    pdf.CellFormat(0, 10, "Phone: "+cv.Phone, "", 1, "L", false, 0, "")
    pdf.CellFormat(0, 10, "LinkedIn: "+cv.Linkedin, "", 1, "L", false, 0, "")
    pdf.Ln(10)

    pdf.CellFormat(0, 10, "Objective", "", 1, "L", false, 0, "")
    pdf.MultiCell(0, 10, cv.Objective, "", "L", false)
    pdf.Ln(10)

    pdf.CellFormat(0, 10, "Work Experience", "", 1, "L", false, 0, "")
    pdf.MultiCell(0, 10, cv.WorkExperience, "", "L", false)
    pdf.Ln(10)

    pdf.CellFormat(0, 10, "Organization Experience", "", 1, "L", false, 0, "")
    pdf.MultiCell(0, 10, cv.OrganizationExperience, "", "L", false)
    pdf.Ln(10)

    pdf.CellFormat(0, 10, "Education", "", 1, "L", false, 0, "")
    pdf.MultiCell(0, 10, cv.Education, "", "L", false)
    pdf.Ln(10)

    pdf.CellFormat(0, 10, "Certification", "", 1, "L", false, 0, "")
    pdf.MultiCell(0, 10, cv.Certification, "", "L", false)
    pdf.Ln(10)

    pdf.CellFormat(0, 10, "Achievement", "", 1, "L", false, 0, "")
    pdf.MultiCell(0, 10, cv.Achievement, "", "L", false)

    var buf bytes.Buffer
    err := pdf.Output(&buf)
    if err != nil {
        return nil, err
    }

    return buf.Bytes(), nil
}
