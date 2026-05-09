package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/th"
)

type ConnectionList struct {
	widget.BaseWidget

	connections []*ConnectionButton

	selected int

	ChangeView	func(index int)
}

func NewConnectionList(changeView func(int), items []*ConnectionButton) *ConnectionList {
	c := &ConnectionList{
		connections: items,
		selected: -1,
		ChangeView: changeView,
	}
	c.ExtendBaseWidget(c)
	return c
}

func (c *ConnectionList) CreateRenderer() fyne.WidgetRenderer {
	c.ExtendBaseWidget(c)
	t := c.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	background := canvas.NewRectangle(t.Color(theme.ColorNameButton, v))
	background.CornerRadius = t.Size(theme.SizeNameInputRadius)

	header := canvas.NewText("CONNECTIONS", th.Palette.SecondaryText)
	header.TextSize = 10

	objects := []fyne.CanvasObject{
		background,
		header,
	}
	for i, obj := range c.connections {
		obj.OnTapped = func(pe *fyne.PointEvent) {
			c.SetSelected(i)
		}
		objects = append(objects, obj)
	}

	r := &connectionListRenderer{
		BaseRenderer: NewBaseRenderer(objects),
		header: header,
		connections: c.connections,
		background: background,
	}
	r.applyTheme()
	return r
}

func (c *ConnectionList) SetSelected(index int) {
	if c.selected != -1 {
		c.connections[c.selected].focused = false
		c.connections[c.selected].Refresh()
	}
	c.selected = index
	c.connections[index].focused = true
	c.connections[index].Refresh()
	if c.ChangeView != nil {
	c.ChangeView(index)
	}
	c.Refresh()
}

type connectionListRenderer struct {
	BaseRenderer

	background *canvas.Rectangle
	header				*canvas.Text
	connections			[]*ConnectionButton
	layout     fyne.Layout
}

func (r *connectionListRenderer) MinSize() fyne.Size {
	hs := r.header.MinSize()

	h := hs.Height + 20
	l := len(r.connections)
	if l < 1 {
		return fyne.NewSize(0, h)
	}

	h = h + float32(r.connections[0].MinSize().Height * 4)
	return fyne.NewSize(0, h)
}

func (r *connectionListRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)

	h := r.header
	hs := fyne.NewSize(size.Width, h.MinSize().Height + 20)
	hp := fyne.NewPos(10, (hs.Height - h.MinSize().Height) / 2)
	h.Move(hp)

	ch := float32(hs.Height)
	for _, obj := range r.connections {
		s := obj.MinSize()
		obj.Resize(fyne.NewSize(size.Width - 2, s.Height))
		obj.Move(fyne.NewPos(1, ch))
		ch = ch + s.Height
	}
}

func (r *connectionListRenderer) applyTheme() {
	t := fyne.CurrentApp().Settings().Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	if bg := r.background; bg != nil {
		bg.StrokeColor = t.Color(theme.ColorNameSeparator, v)
		bg.StrokeWidth = 1
		bg.CornerRadius = 0
	}
}

func (r *connectionListRenderer) Refresh() {
	r.applyTheme()
}

