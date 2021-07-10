package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	app.Settings().SetTheme(&akakunTheme{})
	window := app.NewWindow("akakun")
	window.Resize(fyne.NewSize(800, 500))
	hello := widget.NewLabel("Hello akakun!")
	window.SetContent(container.NewVBox(
		hello,
		widget.NewButton("osu", func() {
			hello.SetText("あかくんだよ")
		}),
	))

	window.ShowAndRun()
}
