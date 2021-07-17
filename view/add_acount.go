package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"github.com/ha-ya4/akakun/component"
	"github.com/ha-ya4/akakun/lib"
)

func AddAccountView(w fyne.Window, data *lib.AkakunDataContainer) fyne.CanvasObject {
	return component.BaseComponent(widget.NewLabel(""))
}
