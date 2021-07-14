package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BaseComponent(content fyne.CanvasObject) fyne.CanvasObject {
	passwordArea := widget.NewEntry()
	passwordArea.PlaceHolder = "passwaord"

	return container.NewVBox(
		passwordArea,
		widget.NewSeparator(),
		content,
	)
}
