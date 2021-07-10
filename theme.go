package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

/*
// 日本語が正しく表示されない問題を解決するためにオリジナルthemeを作成しアプリにセットする必要があるよう
//
// go install fyne.io/fyne/v2/cmd/fyne@<version>でfyneのcliをインストール
// このアプリを例にすると
// fontを用意する
// fyne bundle font/mplus-1c-regular.ttf > bundle.go
// fyne bundle -append font/mplus-1c-bold.ttf >> bundle.go
// 上２つのコマンドを使用してbundleファイル作成
// theme.goを作成し、fyne.Themeインターフェースを満たす構造体を作成する
// 今回はフォント以外はデフォルトのものを使用する
*/

type akakunTheme struct{}

var _ fyne.Theme = (*akakunTheme)(nil)

func (*akakunTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return resourceMplus1cBoldTtf
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return resourceMplus1cRegularTtf
}

func (*akakunTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(n, v)
}

func (*akakunTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*akakunTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}
