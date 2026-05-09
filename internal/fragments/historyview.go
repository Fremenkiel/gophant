package fragments

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/interfaces"
)

func NewHistoryView(w fyne.Window, r interfaces.ErrorReporter) *container.Scroll {
	c := container.NewHScroll(widget.NewLabel("History"))
	c.Hide()
	return c
}

