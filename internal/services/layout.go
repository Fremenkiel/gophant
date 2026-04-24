package services

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type LayoutService struct {
	Sidebar *SidebarService
}

func NewLayoutService() *LayoutService {
	return &LayoutService{Sidebar: NewSidebarService()}
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


