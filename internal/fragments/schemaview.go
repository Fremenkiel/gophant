package fragments

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/Fremenkiel/gophant/v2/internal/interfaces"
	"github.com/Fremenkiel/gophant/v2/internal/menus"
)

func NewSchemaView(w fyne.Window, r interfaces.ErrorReporter) *container.Scroll {
	//acd := dialogs.NewAddConnectionDialog(r, l)
	//ks.MapDefaultKeyBindings(acd.Window)
	cm := menus.NewConnectionMenu(r, w)

	return container.NewHScroll(NewConnectionList(r, cm))
}

