package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"github.com/ha-ya4/akakun/component"
)

func ExportGroupView(w fyne.Window) fyne.CanvasObject {
	return component.BaseComponent(widget.NewLabel("export"))
}
