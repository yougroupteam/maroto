package main

func main() {
	m := maroto{}

	m.Add(Row(
		Col(Text("AnyText", nil)),
	))

	m.Add(Row(
		Col(Text("AnyText", nil),
			Image("asdasd", nil)),
		Col(Text("AnyText", nil)),
	))

	m.PrintNodes()
}
