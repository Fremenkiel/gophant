package containers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/layouts"
)

type BorderBox struct {
	widget.BaseWidget

	elements	[]fyne.CanvasObject
}

func NewBorderBox(objects ...fyne.CanvasObject) *fyne.Container {
	return container.New(layouts.NewBorderBox(), objects...)
}

