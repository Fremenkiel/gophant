package fragments

import (
	"fyne.io/fyne/v2"
	"github.com/Fremenkiel/gophant/v2/internal/containers"
	"github.com/Fremenkiel/gophant/v2/internal/menus"
)

func NewSchemaView(w fyne.Window) *fyne.Container {
	//acd := dialogs.NewAddConnectionDialog(r, l)
	//ks.MapDefaultKeyBindings(acd.Window)
	cm := menus.NewConnectionMenu(w)

	cl := NewConnectionList(cm)
	cs := containers.NewBorderBox(cl)

	return cs
}

