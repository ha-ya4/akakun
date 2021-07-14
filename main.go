package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	"github.com/ha-ya4/akakun/component"
)

func main() {
	app := app.New()
	app.Settings().SetTheme(&akakunTheme{})

	window := app.NewWindow("akakun")
	window.Resize(fyne.NewSize(800, 600))

	menu := createMainMenu(window)
	window.SetMainMenu(menu)

	window.SetContent(component.BaseComponent(widget.NewLabel("content")))
	window.ShowAndRun()
}
