package mml

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"strings"
)

type Interpreter interface {
	GeneratePdfAndSave(mml string, filePathName string) error
	GeneratePdf(mml string) (bytes.Buffer, error)
}

type interpreter struct {

}

func NewInterpreter() *interpreter {
	return &interpreter{}
}

func (s *interpreter) GeneratePdfAndSave(mml string, filePathName string) error {
	doc, err := html.Parse(strings.NewReader(mml))
	if err != nil {
		log.Fatal(err)
	}

	//m := pdf.NewMaroto(consts.Portrait, consts.A4)

	rows := s.getRows(doc)
	for _, row := range rows {
		heightString := s.getValueAttribute(row, "height")
		fmt.Println(heightString)

		cols := s.getCols(row)
		for _, col := range cols {

			components := s.getComponents(col)
			for _, component := range components {
				fmt.Println(component.Attr)
			}
		}

		fmt.Println("")
	}

	return nil
}

func (s *interpreter) GeneratePdf(mml string) (bytes.Buffer, error) {
	panic("implement me")
}

func (s *interpreter) getRows(n *html.Node) []*html.Node {
	rows := []*html.Node{}

	for _, attr := range n.Attr {
		if attr.Key == "class" && strings.Contains(attr.Val, "row") {
			rows = append(rows, n)
			break
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		otherRows := s.getRows(c)
		rows = append(rows, otherRows...)
	}

	return rows
}

func (s *interpreter) getCols(n *html.Node) []*html.Node {
	cols := []*html.Node{}

	for _, attr := range n.Attr {
		if attr.Key == "class" && strings.Contains(attr.Val, "col") {
			cols = append(cols, n)
			break
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		otherCols := s.getCols(c)
		cols = append(cols, otherCols...)
	}

	return cols
}

func (s *interpreter) getComponents(n *html.Node) []*html.Node {
	components := []*html.Node{}

	for _, attr := range n.Attr {
		if attr.Key == "type" {
			components = append(components, n)
			break
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		otherComponents := s.getComponents(c)
		components = append(components, otherComponents...)
	}

	return components
}

func (s *interpreter) getValueAttribute(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}

	return ""
}