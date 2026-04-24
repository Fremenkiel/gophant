package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/Fremenkiel/gophant/v2/internal/services"
)

func main() {
	a := app.New()
	w := a.NewWindow("Main page")
	w.Resize(fyne.NewSize(500, 400))

	closeKey := desktop.CustomShortcut{KeyName: fyne.KeyW, Modifier: fyne.KeyModifierShortcutDefault }
	w.Canvas().AddShortcut(&closeKey, func(shortcut fyne.Shortcut) {
		a.Quit()
	})

	ls := services.NewLayoutService()
	w.SetContent(ls.BuildLayout())
	
	w.ShowAndRun()
}


