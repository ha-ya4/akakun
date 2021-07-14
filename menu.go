package main

import (
	"fyne.io/fyne/v2"
)

func createMenu(w fyne.Window) *fyne.Menu {
	new := fyne.NewMenuItem("new", func() {})
	open := fyne.NewMenuItem("open", func() {})

	return fyne.NewMenu(
		"menu",
		new,
		open,
	)
}

func createAccountMenu(w fyne.Window) *fyne.Menu {
	account := fyne.NewMenuItem("account", func() {})
	impt := fyne.NewMenuItem("import", func() {})
	export := fyne.NewMenuItem("export", func() {})

	return fyne.NewMenu(
		"account",
		account,
		impt,
		export,
	)
}

func createMainMenu(w fyne.Window) *fyne.MainMenu {
	menu := []*fyne.Menu{
		createMenu(w),
		createAccountMenu(w),
	}
	return fyne.NewMainMenu(menu...)
}
