package main

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// 書いてある記述は全て完璧に理解する
// cdata is data structure.
type cdata struct {
	// 最後の演算結果を保管する
	mem int
	// 最後に押した演算キーを保管する
	cal string
	// 演算直後かどうかを示すフラグ
	flg bool
}

// createNumButtons create number buttons.
// 引数に関数を使う関数 -> 高階関数
// 高階関数の場合、型は func (引数) 戻り値 のように指定する
func createNumButtons(f func(v int)) *fyne.Container {
	// 第一引数に用意したレイアウトを使ってウィジェットを配置するコンテナを作成する
	// fyneパッケージのContainer構造体が作成される（正確にはそのポインタ）
	c := fyne.NewContainerWithLayout(
		layout.NewGridLayout(3),
		// 数字のボタンを用意
		// 第二引数にボタンを押した際の処理を指定（外部の関数を組み込む）
		widget.NewButton(strconv.Itoa(7), func() { f(7) }),
		widget.NewButton(strconv.Itoa(8), func() { f(8) }),
		widget.NewButton(strconv.Itoa(9), func() { f(9) }),
		widget.NewButton(strconv.Itoa(4), func() { f(4) }),
		widget.NewButton(strconv.Itoa(5), func() { f(5) }),
		widget.NewButton(strconv.Itoa(6), func() { f(6) }),
		widget.NewButton(strconv.Itoa(1), func() { f(1) }),
		widget.NewButton(strconv.Itoa(2), func() { f(2) }),
		widget.NewButton(strconv.Itoa(3), func() { f(3) }),
		widget.NewButton(strconv.Itoa(0), func() { f(0) }),
	)
	return c
}

// creteCalcButtons create operation-symbol button.
// 演算キーを設定する
func createCalcButtons(f func(c string)) *fyne.Container {
	c := fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		// ボタンを生成
		widget.NewButton("CL", func() {
			f("CL")
		}),
		widget.NewButton("/", func() {
			f("/")
		}),
		widget.NewButton("*", func() {
			f("*")
		}),
		widget.NewButton("+", func() {
			f("+")
		}),
		widget.NewButton("-", func() {
			f("-")
		}),
	)
	return c
}

// main function
func main() {

	a := app.New()
	w := a.NewWindow("calculator")
	// 固定サイズウインドウにする
	w.SetFixedSize(true)
	// 入力した数字を表示する
	l := widget.NewLabel("0")
	// 文字揃えを表す変数（テキストの終わり位置に揃える（通常は右揃え））
	l.Alignment = fyne.TextAlignTrailing

	// 構造体を生成して変数に格納
	data := cdata{
		mem: 0,
		cal: "",
		flg: false,
	}

	// calc is calculate.
	// nは新しく入力された文字
	calc := func(n int) {
		// data.calは演算記号を保存する
		switch data.cal {
		case "":
			// data.memは演算結果を保存する
			data.mem = n
		case "+":
			data.mem += n
		case "-":
			data.mem -= n
		case "*":
			data.mem *= n
		case "/":
			data.mem /= n
		}
		l.SetText(strconv.Itoa(data.mem))
		data.flg = true
	}

	// pushNum is number button action
	// 数字キーを押した際に呼び出される処理
	pushNum := func(v int) {
		s := l.Text
		if data.flg {
			s = "0"
			data.flg = false
		}
		s += strconv.Itoa(v)
		n, err := strconv.Atoi(s)
		if err == nil {
			l.SetText(strconv.Itoa(n))
		}
	}

	// pushCalc is operatoin symbol button action.
	// 演算キーを押したときの処理
	// cは押したキーの記号
	pushCalc := func(c string) {
		if c == "CL" {
			l.SetText("0")
			data.mem = 0
			data.flg = false
			data.cal = ""
			return
		}
		n, er := strconv.Atoi(l.Text)
		if er != nil {
			return
		}
		// 演算し、最後に押した演算キーの値を更新する
		calc(n)
		data.cal = c
	}

	// pushEnter is enter button action.
	// エンターボタンを押した時の処理
	pushEnter := func() {
		n, er := strconv.Atoi(l.Text)
		if er != nil {
			return
		}
		calc(n)
		data.cal = ""
	}

	// ここから理解する
	k := createNumButtons(pushNum)
	c := createCalcButtons(pushCalc)
	e := widget.NewButton("Enter", pushEnter)

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(
				l, e, nil, c,
			),
			l, e, k, c,
		),
	)
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
