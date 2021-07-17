package view

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/ha-ya4/akakun/component"
	"github.com/ha-ya4/akakun/lib"
	"github.com/syndtr/goleveldb/leveldb"
)

func OpenGroupView(w fyne.Window, data *lib.AkakunDataContainer) fyne.CanvasObject {
	groups := []fyne.CanvasObject{}
	for _, group := range data.Groups {
		button := createOpenGroupButton(group, data, w)
		groups = append(groups, button)
	}

	if len(groups) == 0 {
		dialog.ShowInformation("", "グループが登録されていません", w)
	}

	return component.BaseComponent(
		groups...,
	)
}

func createOpenGroupButton(g lib.AkakunAccount, data *lib.AkakunDataContainer, w fyne.Window) *widget.Button {
	return widget.NewButton(g.Name, func() {
		// g.Pathが最後のスラッシュなしで記録されるのでgroup名とpathの間にいれる
		path := fmt.Sprintf("%s/%s%s", g.Path, g.Name, lib.DBSuffix)
		db, err := leveldb.OpenFile(path, nil)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		// すでに別のグループが開かれている可能性があるのでcloseを試みる
		if err = data.CloseDB(); err != nil {
			dialog.ShowError(err, w)
			db.Close()
			return
		}
		data.DB = db
		// 指定されたグループのアカウント一覧へページ偏移する
		w.SetContent(GroupView(w, g.Name))
	})
}
