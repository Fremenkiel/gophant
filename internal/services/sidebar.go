package services

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/dialog"
	"github.com/Fremenkiel/gophant/v2/internal/layouts"
	"github.com/Fremenkiel/gophant/v2/internal/models"
	"github.com/Jipok/go-persist"
)

type SidebarService struct {
	AddConnectionDialog		*dialog.AddConnectionDialog
}

func NewSidebarService(acd *dialog.AddConnectionDialog) *SidebarService {
	return &SidebarService{AddConnectionDialog: acd}
}

func (s *SidebarService) BuildSidebar() *fyne.Container {
	data := createSidebarElements()

	l := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(data[lii].Name)
		},
		)

	l.OnSelected = func(i widget.ListItemID) {
		connections, err := persist.OpenSingleMap[models.Connection]("connections.db")
		if err != nil {
			log.Fatal(err)
		}
		defer connections.Store.Close()

		id := data[i].ID
		log.Print(id)

		key := ""
		if id != nil {
			key = id.String()
		}
		log.Print(key)

		connections.Delete(key)
		data = createSidebarElements()
		l.UnselectAll()
		l.Refresh()
	}

	scroll := container.NewVScroll(l)

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			s.AddConnectionDialog.Open()
			/*
			connections, err := persist.OpenSingleMap[models.Connection]("connections.db")
			if err != nil {
				log.Fatal(err)
			}
			defer connections.Store.Close()

			id, err := uuid.NewV7()
			if err != nil {
				log.Fatal(err)
			}
			log.Print("Add val")
			connections.Set(id.String(), models.Connection{
				ID: &id,
				Name: "DB1",
			})
			data = createSidebarElements()
			l.Refresh()
			*/
		}),
		)


	return container.New(&layouts.Sidebar{},
		toolbar,
		scroll,
		)
}

func createSidebarElements() []models.Connection {
	databases, err := persist.OpenSingleMap[models.Connection]("connections.db")
	if err != nil {
		log.Fatal(err)
	}
	defer databases.Store.Close()

	var connections []models.Connection
	databases.Range(func(key string, value models.Connection) bool {
		connections = append(connections, value)
		return true
	})
	return connections
}
