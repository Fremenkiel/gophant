package menus

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/models"
	"github.com/Jipok/go-persist"
	"github.com/google/uuid"
)

type ConnectionMenu struct {
	Window					fyne.Window
	PopUp					*widget.PopUpMenu
}

func NewConnectionMenu(a fyne.App, w fyne.Window) *ConnectionMenu {
	return &ConnectionMenu{Window: w}
}


func (d *ConnectionMenu) Open(pos fyne.Position, u uuid.UUID, refresh func()) {
	m := fyne.NewMenu("Test", fyne.NewMenuItem("Remove Connection", func() {
		connections, err := persist.OpenSingleMap[models.Connection]("connections.db")
		if err != nil {
			log.Fatal(err)
		}
		defer connections.Store.Close()

		connections.Delete(u.String())
		refresh()
	}))
	p := widget.NewPopUpMenu(m, d.Window.Canvas())
	p.Resize(fyne.NewSize(300, 200))
	p.Move(pos)
	p.Show()

	d.PopUp = p
}
