package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Fremenkiel/gophant/v2/internal/services"
	"github.com/Fremenkiel/gophant/v2/internal/theme"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&theme.GophantTheme{})

	ks := services.NewKeyShortcutService()
	ls := services.NewLayoutService(a, ks)

	w := a.NewWindow("Main page")
	w.Resize(fyne.NewSize(500, 400))
	w.SetContent(ls.BuildLayout())
	ls.KeyShortcut.MapDefaultKeyBindings(w)
	w.Show()
	
	a.Run()
}


