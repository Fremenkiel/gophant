package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type SelectList struct {
	widget.List
}

func NewSelectList(length func() int, createItem func() fyne.CanvasObject, updateItem func(widget.ListItemID, fyne.CanvasObject)) *SelectList {
	b := &SelectList{*widget.NewList(length, createItem, updateItem)}
	return b
}

func (l *SelectList) IsSelected(id widget.ListItemID) {
}

