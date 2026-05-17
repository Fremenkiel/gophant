package fragments

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/Fremenkiel/gophant/v2/internal/containers"
	"github.com/Fremenkiel/gophant/v2/internal/elements"
	"github.com/Fremenkiel/gophant/v2/internal/fs"
	"github.com/Fremenkiel/gophant/v2/internal/th"
)

func NewSidebar(w fyne.Window) *fyne.Container {
	sc := NewSchemaView(w,)
	qu := NewQueryView(w)
	hi := NewHistoryView(w)

	toolbar := containers.NewSidebarTab(
		func(i int) {
			sc.Hide()
			qu.Hide()
			hi.Hide()

			switch i {
			case 0:
				sc.Show()
			case 1:
				qu.Show()
			case 2:
				hi.Show()
			}
		},
		elements.NewSidebarTab("Schema", fs.IconNameDB, nil),
		elements.NewSidebarTab("Queries", fs.IconNameBookmark, nil),
		elements.NewSidebarTab("History", fs.IconNameHistory, nil),
		)


	c := container.New(&th.Sidebar{},
		toolbar,
		sc,
		qu,
		hi,
		)

	return c
}

