package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

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

	w.SetContent(container.NewVBox(
		form,
		hello,
		))
	
	w.ShowAndRun()
}


