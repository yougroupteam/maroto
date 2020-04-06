package main

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/mml"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	interpreter := mml.NewInterpreter()

	htmlBytes, errRead := ioutil.ReadFile("internal/examples/barcodegrid/mml/index.html")
	if errRead != nil {
		fmt.Println("Got error while opening file:", errRead)
		os.Exit(1)
	}

	htmlString := string(htmlBytes)
	htmlString = strings.Replace(htmlString, "\n", "", -1)
	htmlString = strings.Replace(htmlString, "\t", "", -1)

	err := interpreter.GeneratePdfAndSave(htmlString, "internal/examples/pdfs/barcodegrid_mml.pdf")
	if err != nil {
		panic(err)
	}
}
