package fragments

import (
	"fyne.io/fyne/v2"
	"github.com/Fremenkiel/gophant/v2/internal/containers"
	"github.com/Fremenkiel/gophant/v2/internal/interfaces"
	"github.com/Fremenkiel/gophant/v2/internal/menus"
)

func NewSchemaView(w fyne.Window, r interfaces.ErrorReporter) *fyne.Container {
	//acd := dialogs.NewAddConnectionDialog(r, l)
	//ks.MapDefaultKeyBindings(acd.Window)
	cm := menus.NewConnectionMenu(r, w)

	cl := NewConnectionList(r, cm)
	cs := containers.NewBorderBox(cl)

	return cs
}

