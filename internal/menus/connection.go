package menus

import (
	"fyne.io/fyne/v2"
)

type ConnectionMenu struct {
	Window					fyne.Window
}

func NewConnectionMenu(a fyne.App) *ConnectionMenu {
	w := a.NewWindow("Connection menu")
	w.Resize(fyne.NewSize(500, 400))

	return &ConnectionMenu{Window: w}
}


func (d *ConnectionMenu) Open() {
	d.Window.Show()
}
