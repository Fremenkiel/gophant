package menus

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ConnectionMenu struct {
	PopUp					*widget.PopUpMenu
}

func NewConnectionMenu(a fyne.App, w fyne.Window) *ConnectionMenu {
	m := fyne.NewMenu("Test", fyne.NewMenuItem("First", func() {}))
	p := widget.NewPopUpMenu(m, w.Canvas())
	p.Resize(fyne.NewSize(300, 200))

	return &ConnectionMenu{PopUp: p}
}


func (d *ConnectionMenu) Open(pos fyne.Position) {
	d.PopUp.Move(pos)
	d.PopUp.Show()
}
