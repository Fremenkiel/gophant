package fragments

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewQueryView(w fyne.Window) *container.Scroll {
	c := container.NewHScroll(widget.NewLabel("Query"))
	c.Hide()
	return c
}

