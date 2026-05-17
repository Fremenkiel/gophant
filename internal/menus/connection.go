package menus

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/handlers"
	"github.com/Fremenkiel/gophant/v2/internal/models"
	"github.com/Fremenkiel/gophant/v2/internal/repositories"
)

type ConnectionMenu struct {
	Window					fyne.Window
	PopUp					*widget.PopUpMenu
	repo					*repositories.ConnectionRepository
}

func NewConnectionMenu(w fyne.Window) *ConnectionMenu {
	return &ConnectionMenu{Window: w, repo: repositories.NewConnectionRepository()}
}


func (d *ConnectionMenu) Open(pos fyne.Position, c *handlers.ConnectionHandler, refresh, reload func()) {
	i1 := fyne.NewMenuItem("Remove Connection", func() {
		d.repo.Delete(*c.Connection)
		reload()
	})

	i2 := fyne.NewMenuItem("Disconnect", func() {
		c.Disconnect()
		d.repo.Update(*c.Connection)
		refresh()
	})
	if c.Connection.Status == models.OFFLINE {
		i2 = fyne.NewMenuItem("Connect", func() {
			c.Connect()
			d.repo.Create(*c.Connection)
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
