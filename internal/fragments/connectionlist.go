package fragments

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/elements"
	"github.com/Fremenkiel/gophant/v2/internal/handlers"
	"github.com/Fremenkiel/gophant/v2/internal/menus"
	"github.com/Fremenkiel/gophant/v2/internal/models"
	"github.com/Fremenkiel/gophant/v2/internal/themes"
	"github.com/Jipok/go-persist"
)

type ConnectionList struct {
	List	*widget.List
	Data	[]*handlers.ConnectionHandler
}

func NewConnectionList(a fyne.App, cm *menus.ConnectionMenu) *ConnectionList {
	cl := &ConnectionList{Data: createSidebarElements()}

	var l *widget.List
	l = widget.NewList(
		func() int {
			return len(cl.Data)
		},
		func() fyne.CanvasObject {
			c := canvas.NewCircle(themes.Palette.Background)
			return elements.NewIconBox("Template", c, nil, nil, nil)
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			h := cl.Data[lii]
			d := h.Connection
			lbl := co.(*elements.IconBox)
			sc := themes.Palette.Danger
			if d.Status == models.ONLINE {
				sc = themes.Palette.Success
			}

			c := canvas.NewCircle(sc)
			lbl.SetContent(d.Name, c)
			lbl.OnTapped = func(pe *fyne.PointEvent) {
				l.Select(lii)
			}
			lbl.OnTappedSecondary = func(pe *fyne.PointEvent) {
				cm.Open(pe.AbsolutePosition, h, cl.Refresh, cl.Reload)
			}
			lbl.OnDoubleTapped = func (pe *fyne.PointEvent) {
				log.Print("DT")
			}
		},
	)

	l.OnSelected = func(i widget.ListItemID) {
	}

	cl.List = l

	return cl
}

func (c *ConnectionList) Refresh() {
	c.List.UnselectAll()
	c.List.Refresh()
}

func (c *ConnectionList) Reload() {
	c.Data = createSidebarElements()
	c.List.UnselectAll()
	c.List.Refresh()
}

func createSidebarElements() []*handlers.ConnectionHandler {
	databases, err := persist.OpenSingleMap[models.Connection]("connections.db")
	if err != nil {
		log.Fatal(err)
	}
	defer databases.Store.Close()

	var connections []*handlers.ConnectionHandler
	databases.Range(func(k string, v models.Connection) bool {
		connections = append(connections, handlers.NewConnectionHandler(&v))
		return true
	})
	return connections
}
