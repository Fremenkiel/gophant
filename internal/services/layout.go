package services

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/dialog"
)

type LayoutService struct {
	Sidebar 		*SidebarService
	KeyShortcut	*KeyShortcutService
}

func NewLayoutService(a fyne.App, ks *KeyShortcutService) *LayoutService {
	acd := dialog.NewAddConnectionDialog(a)
	ks.MapDefaultKeyBindings(acd.Window)

	return &LayoutService{Sidebar: NewSidebarService(acd), KeyShortcut: ks}
}

func (s *LayoutService) BuildLayout() *container.Split {
	return container.NewHSplit(
		s.Sidebar.BuildSidebar(),
		s.buildMainLayout(),
		)
}

func (s *LayoutService) buildMainLayout() *fyne.Container {
	textArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Query",Widget: textArea},
		},
	}

	return container.NewVBox(
		form,
		)
}


