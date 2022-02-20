package main

import (
	// アプリケーションとウインドウを作成するのに使う
	"fyne.io/fyne/app"
	// ウインドウ内に配置する部品を作成するために使う
	"fyne.io/fyne/widget"
)

func main() {
	// 新しいアプリケーションを作成
	a := app.New()

	// 新しいウインドウを作成
	w := a.NewWindow("Hello")
	// ウインドウに表示するコンテンツを設定
	w.SetContent(
		// これで複数のコンテンツを表示できる
		widget.NewHBox(
			widget.NewLabel("Hello Fyne!"),
			widget.NewLabel("This is sample application!"),
		),
	)

	// ウインドウを表示し実行する
	w.ShowAndRun()
}
