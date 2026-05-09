package fragments

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/interfaces"
)

type MainLayout struct {
	Sidebar 		*fyne.Container
}

func NewMainLayout(w fyne.Window, r interfaces.ErrorReporter) *MainLayout {

	return &MainLayout{Sidebar: NewSidebar(w, r)}
}

func (s *MainLayout) BuildLayout() *container.Split {
	c := container.NewHSplit(
		s.Sidebar,
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


