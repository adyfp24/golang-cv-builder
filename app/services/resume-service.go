package services

import (
	"bytes"
	"github.com/adyfp24/golang-cv-builder/app/models"
	"github.com/jung-kurt/gofpdf"
)

func GeneratePDF(cv models.Resume) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Times", "B", 17)
	pdf.CellFormat(0, 10, cv.Fullname, "", 1, "C", false, 0, "")

	pdf.SetFont("Times", "", 12)
	pdf.CellFormat(0, 5, cv.Linkedin+" | "+cv.Phone+" | "+cv.Email, "", 1, "C", false, 0, "")
	pdf.CellFormat(0, 5, cv.Portofolio+" | "+cv.City, "", 1, "C", false, 0, "")
	pdf.Ln(5)

	pdf.SetFont("Times", "B", 12)
	pdf.CellFormat(0, 10, "OBJECTIVE", "", 1, "L", false, 0, "")
    pdf.Line(10, pdf.GetY(), 200, pdf.GetY())
    pdf.Ln(2)
	pdf.SetFont("Times", "", 12)
	pdf.MultiCell(0, 5, cv.Objective, "", "L", false)
	pdf.Ln(5)

	pdf.SetFont("Times", "B", 12)
	pdf.CellFormat(0, 10, "EDUCATION", "", 1, "L", false, 0, "")
    pdf.Line(10, pdf.GetY(), 200, pdf.GetY())
    pdf.Ln(2)
	pdf.SetFont("Times", "", 12)
	pdf.MultiCell(0, 5, cv.Education, "", "L", false)
	pdf.Ln(5)

	pdf.SetFont("Times", "B", 12)
	pdf.CellFormat(0, 10, "SKILLS", "", 1, "L", false, 0, "")
    pdf.Line(10, pdf.GetY(), 200, pdf.GetY())
    pdf.Ln(2)
	pdf.SetFont("Times", "", 12)
	pdf.MultiCell(0, 5, cv.Skill, "", "L", false)
	pdf.Ln(5)

	pdf.SetFont("Times", "B", 12)
	pdf.CellFormat(0, 10, "WORK EXPERIENCE", "", 1, "L", false, 0, "")
    pdf.Line(10, pdf.GetY(), 200, pdf.GetY())
    pdf.Ln(2)
	pdf.SetFont("Times", "", 12)
	pdf.MultiCell(0, 5, cv.WorkExperience, "", "L", false)
	pdf.Ln(5)

	pdf.SetFont("Times", "B", 12)
	pdf.CellFormat(0, 10, "ORGANIZATION EXPERIENCE", "", 1, "L", false, 0, "")
    pdf.Line(10, pdf.GetY(), 200, pdf.GetY())
    pdf.Ln(2)
	pdf.SetFont("Times", "", 12)
	pdf.MultiCell(0, 5, cv.OrganizationExperience, "", "L", false)
	pdf.Ln(5)


	pdf.SetFont("Times", "B", 12)
	pdf.CellFormat(0, 10, "CERTIFICATION", "", 1, "L", false, 0, "")
    pdf.Line(10, pdf.GetY(), 200, pdf.GetY())
    pdf.Ln(2)
	pdf.SetFont("Times", "", 12)
	pdf.MultiCell(0, 5, cv.Certification, "", "L", false)
	pdf.Ln(5)

	pdf.SetFont("Times", "B", 12)
	pdf.CellFormat(0, 10, "ACHIEVEMENT", "", 1, "L", false, 0, "")
    pdf.Line(10, pdf.GetY(), 200, pdf.GetY())
    pdf.Ln(2)
	pdf.SetFont("Times", "", 12)
	pdf.MultiCell(0, 5, cv.Achievement, "", "L", false)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
