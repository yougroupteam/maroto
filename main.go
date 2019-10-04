package main

func main() {
	m := New(Portrait, A4)

	m.Add(Row(20,
		Col(Text("AnyText", nil)),
	))

	m.Add(Row(20,
		Col(Text("AnyText", nil),
			Image("asdasd", nil)),
		Col(Text("AnyText", nil)),
	))

	//m.PrintNodes()
	m.CreatePDF()
}
