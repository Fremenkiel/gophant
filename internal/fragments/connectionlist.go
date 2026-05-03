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
	"github.com/Fremenkiel/gophant/v2/internal/themes"
	"github.com/Jipok/go-persist"
)

type ConnectionList struct {
	List	*widget.List
	Data	[]*handlers.ConnectionHandler

	reporter		interfaces.ErrorReporter
}

func NewConnectionList(a fyne.App, r interfaces.ErrorReporter, cm *menus.ConnectionMenu) *ConnectionList {
	cl := &ConnectionList{Data: createSidebarElements(r), reporter: r}

	cl.List = widget.NewList(
		func() int {
			return len(cl.Data)
		},
		func() fyne.CanvasObject {
			c := canvas.NewCircle(themes.Palette.Background)
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
			sc := themes.Palette.Danger
			if d.Status == models.ONLINE {
				sc = themes.Palette.Success
			}

			c := canvas.NewCircle(sc)
			lbl.SetHeader(d.Name, c,
				nil,
				func (pe *fyne.PointEvent) {
					fetchDatabases(cl, h, lbl)
				},
				func(pe *fyne.PointEvent) {
					cm.Open(pe.AbsolutePosition, h, cl.Refresh, cl.Reload)
				})
			fetchDatabases(cl, h, lbl)
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

func fetchDatabases(cl *ConnectionList, h *handlers.ConnectionHandler, lbl *elements.Collapse) {
	dbs := h.GetDatabases(cl.Refresh)

	var ell []*elements.IconBox
	for _, el := range dbs {
		ell = append(ell, elements.NewIconBox(el.Name, nil, nil, nil, nil))
	}
	lbl.SetContent(ell)
	lbl.Toggle()
}
