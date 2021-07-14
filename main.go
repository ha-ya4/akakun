package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	"github.com/ha-ya4/akakun/component"
	"github.com/ha-ya4/akakun/lib"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()

	app := app.New()
	app.Settings().SetTheme(&akakunTheme{})

	window := app.NewWindow("akakun")
	window.Resize(fyne.NewSize(800, 600))
	defer window.Close()

	data, err := lib.CreateDataContainer()
	if err != nil {
		return
	}

	menu := createMainMenu(window, data)
	window.SetMainMenu(menu)

	window.SetContent(component.BaseComponent(widget.NewLabel("content")))
	window.ShowAndRun()
}
