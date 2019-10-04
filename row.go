package main

import "fmt"

// Row
type row struct {
	height float64
	cols   []col
}

func (s *row) GetHeight() float64 {
	return s.height
}

func (s *row) QtdCols() int {
	return len(s.cols)
}

func Row(height float64, cols ...col) row {
	return row{
		height: height,
		cols:   cols,
	}
}

func (s *row) PrintNodes() {
	fmt.Println("I'm a row")
	for _, col := range s.cols {
		col.PrintNodes()
	}
}
