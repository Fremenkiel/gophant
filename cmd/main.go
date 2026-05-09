package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Fremenkiel/gophant/v2/internal/dialogs"
	"github.com/Fremenkiel/gophant/v2/internal/layouts"
	"github.com/Fremenkiel/gophant/v2/internal/th"
	"github.com/Fremenkiel/gophant/v2/internal/utils"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&th.GophantTheme{})

	w := a.NewWindow("Main page")
	w.Resize(fyne.NewSize(1000, 800))
	w.SetPadded(false)
	w.SetMainMenu(fyne.NewMainMenu(fyne.NewMenu("test", fyne.NewMenuItem("test.action", func() {}))))
	w.SetMaster()

	ed := dialogs.NewErrorDialog()

	ls := layouts.NewMainLayout(w, ed)

	w.SetContent(ls.BuildLayout())
	utils.MapMainKeyBindings(w)

	w.Show()	
	a.Run()
}


