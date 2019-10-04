package main

import "fmt"

// Col
type col struct {
	components []Component
}

func Col(components ...Component) col {
	return col{
		components: components,
	}
}

func (s *col) PrintNodes() {
	fmt.Println("I'm a col")
	for _, compoment := range s.components {
		compoment.PrintNodes()
	}
}
