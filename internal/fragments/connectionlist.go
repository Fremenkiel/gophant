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

func NewConnectionList(a fyne.App) *ConnectionList {
	cm := menus.NewConnectionMenu(a)
	cl := &ConnectionList{Data: createSidebarElements()}

	l := widget.NewList(
		func() int {
			return len(cl.Data)
		},
		func() fyne.CanvasObject {
			l := elements.NewTapSecondaryLabel("Template", func(pe *fyne.PointEvent) {
				log.Print(pe)
				cm.Open()
			})
			return l
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*elements.TapSecondaryLabel).SetText(cl.Data[lii].Name)
		},
		)

	l.OnSelected = func(i widget.ListItemID) {
		connections, err := persist.OpenSingleMap[models.Connection]("connections.db")
		if err != nil {
			log.Fatal(err)
		}
		defer connections.Store.Close()

		id := cl.Data[i].ID
		log.Print(id)

		key := ""
		if id != nil {
			key = id.String()
		}
		log.Print(key)

		connections.Delete(key)
		cl.Refresh()
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
