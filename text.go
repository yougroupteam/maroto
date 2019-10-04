package main

import "fmt"

type text struct {
	content string
	props   *TextProp
}

func (s *text) Apply() {
	fmt.Println("Applying text")
}

func (s *text) PrintNodes() {
	fmt.Println("I'm a text")
}

type TextProp struct {
	// Top is space between the upper cell limit to the barcode, if align is not center
	Top float64
	// Size of the text
	Size float64
	// Extrapolate define if the text will automatically add a new line when
	// text reach the right cell boundary
	Extrapolate bool
}

func Text(content string, textProp *TextProp) Component {
	return &text{
		content: content,
		props:   textProp,
	}
}
