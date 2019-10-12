package main

func main() {
	m := New(Portrait, A4)

	m.Add(Row(20,
		Col(Text("AnyText")),
	))

	m.Add(Row(20,
		Col(Text("AnyText"),
		    Image("image.png")),
		Col(Text("AnyText")),
	))

	m.CreatePDF()
}
