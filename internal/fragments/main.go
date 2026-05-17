package fragments

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MainLayout struct {
	Sidebar 		*fyne.Container
}

func NewMainLayout(w fyne.Window) *MainLayout {

	return &MainLayout{Sidebar: NewSidebar(w)}
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


