package elements

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/models"
	"github.com/Jipok/go-persist"
)

type ConnectionList struct {
	List	*widget.List
	Data	[]models.Connection
}

func NewConnectionList() *ConnectionList {
	cl := &ConnectionList{Data: createSidebarElements()}

	l := widget.NewList(
		func() int {
			return len(cl.Data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(cl.Data[lii].Name)
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
