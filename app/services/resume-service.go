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
    pdf.CellFormat(0, 10, "Name: "+cv.Name, "", 1, "L", false, 0, "")
    pdf.CellFormat(0, 10, "Email: "+cv.Email, "", 1, "L", false, 0, "")
    pdf.CellFormat(0, 10, "Phone: "+cv.Phone, "", 1, "L", false, 0, "")
    pdf.Ln(10)

    pdf.CellFormat(0, 10, "Experience", "", 1, "L", false, 0, "")
    pdf.MultiCell(0, 10, cv.Experience, "", "L", false)
    pdf.Ln(10)
    pdf.CellFormat(0, 10, "Education", "", 1, "L", false, 0, "")
    pdf.MultiCell(0, 10, cv.Education, "", "L", false)

    var buf bytes.Buffer
    err := pdf.Output(&buf)
    if err != nil {
        return nil, err
    }

    return buf.Bytes(), nil
}