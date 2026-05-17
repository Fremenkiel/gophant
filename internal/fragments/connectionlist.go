package fragments

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/dialogs"
	"github.com/Fremenkiel/gophant/v2/internal/elements"
	"github.com/Fremenkiel/gophant/v2/internal/handlers"
	"github.com/Fremenkiel/gophant/v2/internal/menus"
	"github.com/Fremenkiel/gophant/v2/internal/models"
	"github.com/Fremenkiel/gophant/v2/internal/repositories"
)

type ConnectionList struct {
	List	*widget.List
	Data	[]*handlers.ConnectionHandler
}

func NewConnectionList(cm *menus.ConnectionMenu) *elements.ConnectionList {
	connections := createSidebarElements()


	var cb []*elements.ConnectionButton
	for _, obj := range connections {
		cb = append(cb, elements.NewConnectionButton(obj.Connection.Name, obj.Connection.Permission, nil, nil, nil))
	}

	cl := elements.NewConnectionList(cb, nil, nil)

	cl.AddConnection = func(pe *fyne.PointEvent) {
	acd := dialogs.NewAddConnectionDialog(
		func(connection models.Connection) {
			h := handlers.NewConnectionHandler(&connection)
			cb = append(cb, elements.NewConnectionButton(h.Connection.Name, h.Connection.Permission, nil, nil, nil))
			cl.SetContent(cb)
		})
		acd.Open()
	}

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
	repo := repositories.NewConnectionRepository()
	databases, err := repo.GetAll()
	if err != nil {
		dialogs.ReportError(err)
		return nil
	}

	var connections []*handlers.ConnectionHandler

	for _, obj := range databases {
		h := handlers.NewConnectionHandler(&obj)
		connections = append(connections, h)
	}

	return connections
}

func (c *ConnectionList) fetchDatabases(h *handlers.ConnectionHandler, lbl *elements.Collapse) {
	dbs := h.GetDatabases(c.Refresh)

	var ell []*elements.Collapse
	for _, el := range dbs {
		ib := elements.NewIconBox(el.Name, nil,
			func(pe *fyne.PointEvent) {}, 
			nil, nil)
		cl := elements.NewCollapse(ib, nil)
		ib.OnDoubleTapped = func(pe *fyne.PointEvent) {
				c.fetchDatabase(h, cl)
			}
		ell = append(ell, cl)
	}
	lbl.SetContent(ell)
	lbl.Open()
}

func (c *ConnectionList) fetchDatabase(h *handlers.ConnectionHandler, lbl *elements.Collapse) {
}
