package menus

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/handlers"
	"github.com/Fremenkiel/gophant/v2/internal/models"
	"github.com/Jipok/go-persist"
)

type ConnectionMenu struct {
	Window					fyne.Window
	PopUp					*widget.PopUpMenu
}

func NewConnectionMenu(a fyne.App, w fyne.Window) *ConnectionMenu {
	return &ConnectionMenu{Window: w}
}


func (d *ConnectionMenu) Open(pos fyne.Position, c *handlers.ConnectionHandler, refresh, reload func()) {
	i1 := fyne.NewMenuItem("Remove Connection", func() {
		connections, err := persist.OpenSingleMap[models.Connection]("connections.db")
		if err != nil {
			log.Fatal(err)
		}
		defer connections.Store.Close()

		connections.Delete(c.Connection.ID.String())
		reload()
	})

	i2 := fyne.NewMenuItem("Disconnect", func() {
		c.Disconnect()
		connections, err := persist.OpenSingleMap[models.Connection]("connections.db")
		if err != nil {
			log.Fatal(err)
		}
		defer connections.Store.Close()

		connections.Set(c.Connection.ID.String(), *c.Connection)
		refresh()
	})
	if c.Connection.Status == models.OFFLINE {
		i2 = fyne.NewMenuItem("Connect", func() {
			c.Connect()
			connections, err := persist.OpenSingleMap[models.Connection]("connections.db")
			if err != nil {
				log.Fatal(err)
			}
			defer connections.Store.Close()

			connections.Set(c.Connection.ID.String(), *c.Connection)
			refresh()
		})
	}
	m := fyne.NewMenu("Connection Menu", i1, i2)
	p := widget.NewPopUpMenu(m, d.Window.Canvas())
	s := p.Size()
	p.Resize(fyne.NewSize(300, s.Height))
	p.Move(pos)
	p.Show()

	d.PopUp = p
}
