package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

// Maroto
type maroto struct {
	Pdf  gofpdf.Pdf
	rows []row
}

func New(orientation Orientation, pageSize PageSize) maroto {
	fpdf := gofpdf.New(string(orientation), "mm", string(pageSize), "")
	fpdf.SetMargins(10, 10, 10)

	return maroto{
		Pdf: fpdf,
	}
}

func (s *maroto) Add(row row) {
	s.rows = append(s.rows, row)
}

func (s *maroto) PrintNodes() {
	fmt.Println("I'm a maroto")
	for _, row := range s.rows {
		row.PrintNodes()
	}
}

func (s *maroto) CreatePDF() {
	// Recursively build PDF
	rawWidth, _ := s.Pdf.GetPageSize()
	left, _, right, _ := s.Pdf.GetMargins()

	width := rawWidth - left - right

	for index, row := range s.rows {
		fmt.Printf("Row %d, Col Width %f\n", index, width/float64(len(row.cols)))
		for _, col := range row.cols {
			for _, component := range col.components {
				component.Apply()
			}
		}
	}
}
