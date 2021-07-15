package view

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/syndtr/goleveldb/leveldb"

	"github.com/ha-ya4/akakun/component"
	"github.com/ha-ya4/akakun/lib"
)

func NewGroupView(w fyne.Window, data *lib.AkakunDataContainer) fyne.CanvasObject {
	name := widget.NewEntry()
	path := widget.NewEntry()
	formItems := []*widget.FormItem{
		{Text: "group名", Widget: name},
		{Text: "パス", Widget: path},
	}

	form := &widget.Form{
		Items: formItems,
		OnSubmit: func() {
			if err := createNewGroup(name.Text, path.Text, data, w); err != nil {
				dialog.ShowError(err, w)
				return
			}
			// 作成に成功したらformに入力された文字を空にする
			for _, entry := range [...]*widget.Entry{name, path} {
				entry.SetText("")
			}
		},
	}

	return component.BaseComponent(
		container.NewMax(form),
	)
}

func createNewGroup(name, path string, data *lib.AkakunDataContainer, w fyne.Window) error {
	var err error

	// formの入力チェック
	if err = validateNewGroupForm(name, path, data); err != nil {
		return err
	}
	// group作成位置の絶対パスを取得
	abs, err := getNewGroupPathAbs(path, data)
	if err != nil {
		return err
	}
	// group用のDB作成
	if err = createNewGroupDB(name, path, data); err != nil {
		return err
	}
	// db作成に成功したらgroupをの情報を登録する
	if err = data.RegisterGroup(lib.AkakunAccount{Name: name, Path: abs}); err != nil {
		return err
	}

	info := fmt.Sprintf("%sに%sという名前でグループを作成しました", abs, name)
	dialog.ShowInformation("成功", info, w)
	return nil
}

// group名を指定しているか、すでに同じ名前がないか、pathが指定されているかチェック
func validateNewGroupForm(name, path string, data *lib.AkakunDataContainer) error {
	if name == "" {
		return errors.New("group名が指定されていません。")
	}
	// groupの情報から同じ名前のgループがすでに存在していないかチェックする
	for _, g := range data.Group {
		if g.Name == name {
			return errors.New("すでに同じ名前のgroup名が存在します")
		}
	}

	if path == "" {
		return errors.New("group保存先のpathが指定されていません。")
	}

	return nil
}

// groupの保存先をjsonに記録するために保存先の絶対pathを取得する
func getNewGroupPathAbs(path string, data *lib.AkakunDataContainer) (string, error) {
	// 指定されたpathへ移動
	if err := os.Chdir(path); err != nil {
		return "", err
	}

	// 指定されたpathの絶対pathを取得
	p, err := filepath.Abs(".")
	if err != nil {
		return "", err
	}

	// このアプリのあるディレクトリへ戻る
	if err := os.Chdir(data.PrjRoot); err != nil {
		return "", err
	}

	return p, nil
}

func createNewGroupDB(name, path string, data *lib.AkakunDataContainer) error {
	var err error
	if data.DB != nil {
		if err = data.DB.Close(); err != nil {
			return err
		}
	}

	data.DB, err = leveldb.OpenFile(path+name+lib.DBSuffix, nil)
	return err
}
