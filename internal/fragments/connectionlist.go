package fragments

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/elements"
	"github.com/Fremenkiel/gophant/v2/internal/menus"
	"github.com/Fremenkiel/gophant/v2/internal/models"
	"github.com/Jipok/go-persist"
)

type ConnectionList struct {
	List	*widget.List
	Data	[]models.Connection
}

func NewConnectionList(a fyne.App, cm *menus.ConnectionMenu) *ConnectionList {
	cl := &ConnectionList{Data: createSidebarElements()}

	var l *widget.List
	l = widget.NewList(
		func() int {
			return len(cl.Data)
		},
		func() fyne.CanvasObject {
			return elements.NewTapLabel("Template", nil, nil)
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			lbl := co.(*elements.TapLabel)
			lbl.SetText(cl.Data[lii].Name)
			lbl.OnTapped = func() {
				l.Select(lii)
			}
			lbl.OnTappedSecondary = func(pe *fyne.PointEvent) {
				log.Print(pe)
				cm.Open(pe.AbsolutePosition, *cl.Data[lii].ID, cl.Refresh)
			}
		},
	)

	l.OnSelected = func(i widget.ListItemID) {
	}

	cl.List = l

	return cl
}

func (c *ConnectionList) Refresh() {
	c.Data = createSidebarElements()
	c.List.UnselectAll()
	c.List.Refresh()
}

func createSidebarElements() []models.Connection {
	databases, err := persist.OpenSingleMap[models.Connection]("connections.db")
	if err != nil {
		log.Fatal(err)
	}
	defer databases.Store.Close()

	var connections []models.Connection
	databases.Range(func(key string, value models.Connection) bool {
		connections = append(connections, value)
		return true
	})
	return connections
}
