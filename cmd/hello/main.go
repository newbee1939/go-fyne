package main

import (
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	c := 0
	// 新しいアプリケーションを作成
	a := app.New()

	// 新しいウインドウを作成
	w := a.NewWindow("Helloです")
	// 表示するLabelを変数に入れておく
	l := widget.NewLabel("Hello Fyneです!")
	w.SetContent(
		// 複数部品を並べる
		widget.NewVBox(
			l,
			widget.NewButton("Click me!!", func() {
				c++
				// Labelに表示されたテキストを変更する
				l.SetText("count: " + strconv.Itoa(c))
			}),
		),
	)

	// ウインドウを表示し実行する
	w.ShowAndRun()
}
