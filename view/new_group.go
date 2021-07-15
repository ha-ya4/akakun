package view

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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
			createNewGroup(name.Text, path.Text, data)
		},
	}

	return component.BaseComponent(
		container.NewMax(form),
	)
}

func createNewGroup(name, path string, data *lib.AkakunDataContainer) {
	var err error

	if err = validateNewGroupForm(name, path, data); err != nil {
		fmt.Println(err)
	}
	abs, err := getNewGroupPathAbs(path, data)
	if err != nil {
		fmt.Println(err)
	}
	if err = data.RegisterGroup(lib.AkakunAccount{Name: name, Path: abs}); err != nil {
		fmt.Println(err)
	}

	data.DB, err = leveldb.OpenFile(path+name+"_akakun", nil)
	if err != nil {
		fmt.Println(err)
	}
	defer data.DB.Close()
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
