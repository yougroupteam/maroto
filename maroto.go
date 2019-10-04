package main

import "fmt"

// Maroto
type maroto struct {
	rows []row
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

func (s *maroto) Draw() {

}
