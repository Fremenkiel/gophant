package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/layouts"
	"github.com/Fremenkiel/gophant/v2/internal/models"
	"github.com/google/uuid"
)

var databases = []models.Connection{}

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(500, 400))

	closeKey := desktop.CustomShortcut{KeyName: fyne.KeyW, Modifier: fyne.KeyModifierShortcutDefault }
	w.Canvas().AddShortcut(&closeKey, func(shortcut fyne.Shortcut) {
		a.Quit()
	})

	textArea := widget.NewMultiLineEntry()
	hello := widget.NewLabel("")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Query",Widget: textArea},
		},
		OnSubmit: func() {
			hello.SetText(textArea.Text)
		},
	}

	connectionList := widget.NewList(
				func () int {
					return len(databases)
				},
				func () fyne.CanvasObject {
					id, err := uuid.NewV7()
					if err != nil {
						log.Print(err.Error())
					}
					databases = append(databases, models.Connection{ ID: id, Name: "Test"})
					return widget.NewButton("Test", func() {})
				},
				func (i widget.ListItemID, o fyne.CanvasObject) {
					o.(*widget.Button).SetText(databases[i].Name)
				})

	toolbar := widget.NewToolbar(
				widget.NewToolbarAction(theme.ContentAddIcon(), func() {
					connectionList.CreateItem()
				}),
				)

	w.SetContent(container.NewHSplit(
		container.New(&layouts.Sidebar{},
			toolbar,
		container.NewHScroll(
				connectionList,
			)),
		container.NewVBox(
		hello,
		form,
			),
		))
	
	w.ShowAndRun()
}


