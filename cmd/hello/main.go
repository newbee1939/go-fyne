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
		// 3は行数を指定
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
	// 新しいアプリケーションを作成。fyne.Appという値を作成。
	// fyne.Appはアプリケーションの機能を定義したインターフェース。ここにあるメソッドを呼び出すことでアプリケーションを操作する
	// 実際にNewで作成されるのは、fyne.Appインターフェースを実装した構造体の値
	a := app.New()
	// 新しいウインドウを作成
	w := a.NewWindow("calculator")
	// 固定サイズウインドウにする
	w.SetFixedSize(true)
	// 入力した数字を表示する。0は初期値
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
		// 現在ラベルに設定されている値（表示されている値）を取得
		s := l.Text
		// 演算後だった場合はそれぞれ初期化する
		if data.flg {
			s = "0"
			data.flg = false
		}
		// 文字列に変換した上で繋ぎ合せる
		s += strconv.Itoa(v)
		// 数字に変換する
		n, err := strconv.Atoi(s)
		if err == nil {
			// エラーがない場合は文字列にした上でラベルにセットする
			l.SetText(strconv.Itoa(n))
		}
	}

	// ここから！！！！
	// pushCalc is operatoin symbol button action.
	// 演算キーを押したときの処理
	// cは押したキーの記号（+とか*とか）
	pushCalc := func(c string) {
		// clを押した場合は全てを初期化する
		if c == "CL" {
			l.SetText("0")
			data.mem = 0
			data.flg = false
			data.cal = ""
			return
		}
		// 表示されている値を取得
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

	// 電卓の各ボタンを生成する
	// pushNumは各数字のボタンを押したときの動作を定義
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
