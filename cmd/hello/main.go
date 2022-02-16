package main

func main() {
}

func main() {
	a := app.New()

	w := a.NewWindow("Hello")
	w.SetContent(
		widget.NewLabel("Hello Fyne!"),
	)

	w.ShowAndRun()
}
