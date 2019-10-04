package main

import "fmt"

type image struct {
	path  string
	props *TextProp
}

func (s *image) Apply() {

}

func (s *image) PrintNodes() {
	fmt.Println("I'm a text")
}

type ImageProp struct {
	// Top is space between the upper cell limit to the barcode, if align is not center
	Top float64
	// Size of the text
	Size float64
	// Extrapolate define if the text will automatically add a new line when
	// text reach the right cell boundary
	Extrapolate bool
}

func Image(path string, imageProp *ImageProp) Component {
	return &image{
		path:  path,
		props: imageProp,
	}
}
