package main

import "fmt"

// Row
type row struct {
	cols []col
}

func Row(cols ...col) row {
	return row{
		cols: cols,
	}
}

func (s *row) PrintNodes() {
	fmt.Println("I'm a row")
	for _, col := range s.cols {
		col.PrintNodes()
	}
}
