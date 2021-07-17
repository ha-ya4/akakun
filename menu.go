package main

import (
	"fyne.io/fyne/v2"

	"github.com/ha-ya4/akakun/lib"
	"github.com/ha-ya4/akakun/view"
)

func newMenuItem(w fyne.Window, data *lib.AkakunDataContainer) *fyne.MenuItem {
	return fyne.NewMenuItem("new", func() {
		w.SetContent(view.NewGroupView(w, data))
	})
}

func openGroupMenuItem(w fyne.Window, data *lib.AkakunDataContainer) *fyne.MenuItem {
	return fyne.NewMenuItem("open", func() {
		w.SetContent(view.OpenGroupView(w, data))
	})
}

func createMenu(w fyne.Window, data *lib.AkakunDataContainer) *fyne.Menu {
	new := newMenuItem(w, data)
	open := openGroupMenuItem(w, data)

	return fyne.NewMenu(
		"menu",
		new,
		open,
	)
}

func groupMenuItem(w fyne.Window) *fyne.MenuItem {
	return fyne.NewMenuItem("account", func() {
		w.SetContent(view.GroupView(w, ""))
	})
}

func importGroupMenuItem(w fyne.Window) *fyne.MenuItem {
	return fyne.NewMenuItem("import", func() {
		w.SetContent(view.ImportGroupView(w))
	})
}

func exportGroupMenuItem(w fyne.Window) *fyne.MenuItem {
	return fyne.NewMenuItem("export", func() {
		w.SetContent(view.ExportGroupView(w))
	})
}

func createGroupMenu(w fyne.Window) *fyne.Menu {
	account := groupMenuItem(w)
	impt := importGroupMenuItem(w)
	export := exportGroupMenuItem(w)

	return fyne.NewMenu(
		"group",
		account,
		impt,
		export,
	)
}

func createMainMenu(w fyne.Window, data *lib.AkakunDataContainer) *fyne.MainMenu {
	menu := []*fyne.Menu{
		createMenu(w, data),
		createGroupMenu(w),
	}
	return fyne.NewMainMenu(menu...)
}
