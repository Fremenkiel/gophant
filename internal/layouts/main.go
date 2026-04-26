package layouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/dialogs"
	"github.com/Fremenkiel/gophant/v2/internal/fragments"
	"github.com/Fremenkiel/gophant/v2/internal/menus"
	"github.com/Fremenkiel/gophant/v2/internal/utils"
)

type MainLayout struct {
	Sidebar 		*Sidebar
	KeyShortcut	*utils.KeyShortcutUtils
	ConnectionMenu	*menus.ConnectionMenu
}

func NewMainLayout(a fyne.App, w fyne.Window, ks *utils.KeyShortcutUtils) *MainLayout {
	cm := menus.NewConnectionMenu(a, w)
	l := fragments.NewConnectionList(a, cm)
	acd := dialogs.NewAddConnectionDialog(a, l)
	ks.MapDefaultKeyBindings(acd.Window)

	return &MainLayout{Sidebar: NewSidebar(acd, l), KeyShortcut: ks, ConnectionMenu: cm}
}

func (s *MainLayout) BuildLayout() *container.Split {
	c := container.NewHSplit(
		s.Sidebar.BuildSidebar(),
		s.buildMainLayout(),
		)
	c.SetOffset(0.2)
	return c
}

func (s *MainLayout) buildMainLayout() *fyne.Container {
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


