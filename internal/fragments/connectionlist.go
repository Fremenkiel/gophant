package fragments

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/elements"
	"github.com/Fremenkiel/gophant/v2/internal/handlers"
	"github.com/Fremenkiel/gophant/v2/internal/interfaces"
	"github.com/Fremenkiel/gophant/v2/internal/menus"
	"github.com/Fremenkiel/gophant/v2/internal/models"
	"github.com/Fremenkiel/gophant/v2/internal/th"
	"github.com/Jipok/go-persist"
)

type ConnectionList struct {
	List	*widget.List
	Data	[]*handlers.ConnectionHandler

	reporter		interfaces.ErrorReporter
}

func NewConnectionList(r interfaces.ErrorReporter, cm *menus.ConnectionMenu) *ConnectionList {
	cl := &ConnectionList{Data: createSidebarElements(r), reporter: r}

	cl.List = widget.NewList(
		func() int {
			return len(cl.Data)
		},
		func() fyne.CanvasObject {
			c := canvas.NewCircle(th.Palette.Background)
			b := elements.NewIconBox("Header", c, nil, nil, nil)
			return elements.NewCollapse(b, nil)
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			h := cl.Data[lii]
			if h == nil {
				log.Print("No Connection fould")
				return 
			}
			d := h.Connection
			lbl := co.(*elements.Collapse)
			sc := th.Palette.Danger
			if d.Status == models.ONLINE {
				sc = th.Palette.Success
			}

			c := canvas.NewCircle(sc)
			lbl.SetHeader(d.Name, c,
				nil,
				func (pe *fyne.PointEvent) {
					if lbl.Opened { lbl.Close() } else { cl.fetchDatabases(h, lbl) }
				},
				func(pe *fyne.PointEvent) {
					cm.Open(pe.AbsolutePosition, h, cl.Refresh, cl.Reload)
				})
		},
		)

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
