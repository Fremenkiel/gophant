package fragments

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/dialogs"
	"github.com/Fremenkiel/gophant/v2/internal/elements"
	"github.com/Fremenkiel/gophant/v2/internal/handlers"
	"github.com/Fremenkiel/gophant/v2/internal/interfaces"
	"github.com/Fremenkiel/gophant/v2/internal/menus"
	"github.com/Fremenkiel/gophant/v2/internal/models"
	"github.com/Jipok/go-persist"
)

type ConnectionList struct {
	List	*widget.List
	Data	[]*handlers.ConnectionHandler

	reporter		interfaces.ErrorReporter
}

func NewConnectionList(r interfaces.ErrorReporter, cm *menus.ConnectionMenu) *elements.ConnectionList {
	connections := createSidebarElements(r)


	var cb []*elements.ConnectionButton
	for _, obj := range connections {
		cb = append(cb, elements.NewConnectionButton(obj.Connection.Name, obj.Connection.Permission, nil, nil, nil))
	}

	cl := elements.NewConnectionList(cb, nil, nil)
	acd := dialogs.NewAddConnectionDialog(r,
		func(connection models.Connection) {
			h := handlers.NewConnectionHandler(r, &connection)
			cb = append(cb, elements.NewConnectionButton(h.Connection.Name, h.Connection.Permission, nil, nil, nil))
			cl.SetContent(cb)
		})

	cl.AddConnection = func(pe *fyne.PointEvent) {
		acd.Open()
	}

	return cl
}

func (c *ConnectionList) Refresh() {
	c.List.UnselectAll()
	c.List.Refresh()
}

func (c *ConnectionList) Reload() {
	c.Data = createSidebarElements(c.reporter)
	c.List.UnselectAll()
	c.List.Refresh()
}

func createSidebarElements(r interfaces.ErrorReporter) []*handlers.ConnectionHandler {
	databases, err := persist.OpenSingleMap[models.Connection]("connections.db")
	if err != nil {
		r.Report(err)
		return nil
	}
	defer databases.Store.Close()

	var connections []*handlers.ConnectionHandler
	databases.Range(func(k string, v models.Connection) bool {
		h := handlers.NewConnectionHandler(r, &v)
		connections = append(connections, h)
		return true
	})
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
