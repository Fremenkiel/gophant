package fragments

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/dialog"
	"github.com/Fremenkiel/gophant/v2/internal/elements"
	"github.com/Fremenkiel/gophant/v2/internal/themes"
)

type SidebarFragment struct {
	AddConnectionDialog		*dialog.AddConnectionDialog
	ConnectionList				*elements.ConnectionList
}

func NewSidebarFragment(acd *dialog.AddConnectionDialog, cl *elements.ConnectionList) *SidebarFragment {
	return &SidebarFragment{AddConnectionDialog: acd, ConnectionList: cl}
}

func (s *SidebarFragment) BuildSidebar() *fyne.Container {
	scroll := container.NewVScroll(s.ConnectionList.List)

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

