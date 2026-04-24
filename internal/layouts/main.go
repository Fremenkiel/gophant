package layouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/dialogs"
	"github.com/Fremenkiel/gophant/v2/internal/fragments"
	"github.com/Fremenkiel/gophant/v2/internal/utils"
)

type MainLayout struct {
	Sidebar 		*Sidebar
	KeyShortcut	*utils.KeyShortcutUtils
}

func NewMainLayout(a fyne.App, ks *utils.KeyShortcutUtils) *MainLayout {
	l := fragments.NewConnectionList(a)
	acd := dialogs.NewAddConnectionDialog(a, l)
	ks.MapDefaultKeyBindings(acd.Window)

	return &MainLayout{Sidebar: NewSidebar(acd, l), KeyShortcut: ks}
}

func (s *MainLayout) BuildLayout() *container.Split {
	return container.NewHSplit(
		s.Sidebar.BuildSidebar(),
		s.buildMainLayout(),
		)
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


