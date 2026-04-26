package layouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/dialogs"
	"github.com/Fremenkiel/gophant/v2/internal/fragments"
	"github.com/Fremenkiel/gophant/v2/internal/interfaces"
	"github.com/Fremenkiel/gophant/v2/internal/themes"
)

type Sidebar struct {
	AddConnectionDialog		*dialogs.AddConnectionDialog
	ConnectionList				*fragments.ConnectionList

	reporter							interfaces.ErrorReporter
}

func NewSidebar(r interfaces.ErrorReporter, acd *dialogs.AddConnectionDialog, cl *fragments.ConnectionList) *Sidebar {
	return &Sidebar{AddConnectionDialog: acd, ConnectionList: cl, reporter: r}
}

func (s *Sidebar) BuildSidebar() *fyne.Container {
	scroll := container.NewVScroll(container.NewHScroll(s.ConnectionList.List))

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			s.AddConnectionDialog.Open()
		}),
		)


	return container.New(&themes.Sidebar{},
		toolbar,
		scroll,
		)
}

