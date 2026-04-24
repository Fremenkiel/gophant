package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Fremenkiel/gophant/v2/internal/layouts"
	"github.com/Fremenkiel/gophant/v2/internal/themes"
	"github.com/Fremenkiel/gophant/v2/internal/utils"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&themes.GophantTheme{})

	ks := utils.NewKeyShortcutUtils()
	ls := layouts.NewMainLayout(a, ks)

	w := a.NewWindow("Main page")
	w.Resize(fyne.NewSize(500, 400))
	w.SetContent(ls.BuildLayout())
	ls.KeyShortcut.MapDefaultKeyBindings(w)
	w.Show()
	
	a.Run()
}


