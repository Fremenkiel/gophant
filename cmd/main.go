package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Fremenkiel/gophant/v2/internal/dialogs"
	"github.com/Fremenkiel/gophant/v2/internal/layouts"
	"github.com/Fremenkiel/gophant/v2/internal/themes"
	"github.com/Fremenkiel/gophant/v2/internal/utils"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&themes.GophantTheme{})

	w := a.NewWindow("Main page")
	w.Resize(fyne.NewSize(1000, 800))

	ed := dialogs.NewErrorDialog()

	ks := utils.NewKeyShortcutUtils()
	ls := layouts.NewMainLayout(w, ks, ed)

	w.SetContent(ls.BuildLayout())
	ls.KeyShortcut.MapDefaultKeyBindings(w)
	w.Show()
	
	a.Run()
}


