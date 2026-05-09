package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ConnectionList struct {
	widget.BaseWidget

	connections []*ConnectionButton

	selected int

	ChangeView	func(index int)
}

func NewConnectionList(changeView func(int), items []*ConnectionButton) *ConnectionList {
	c := &ConnectionList{connections: items, selected: -1, ChangeView: changeView}
	c.ExtendBaseWidget(c)
	return c
}

func (c *ConnectionList) CreateRenderer() fyne.WidgetRenderer {
	c.ExtendBaseWidget(c)
	t := c.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	background := canvas.NewRectangle(t.Color(theme.ColorNameButton, v))
	background.CornerRadius = t.Size(theme.SizeNameInputRadius)

	objects := []fyne.CanvasObject{
		background,
	}
	for i, obj := range c.connections {
		obj.OnTapped = func(pe *fyne.PointEvent) {
			c.SetSelected(i)
		}
		objects = append(objects, obj)
	}

	r := &connectionListRenderer{
		BaseRenderer: NewBaseRenderer(objects),
		connections: c.connections,
		background: background,
	}
	r.applyTheme()
	return r
}

func (c *ConnectionList) SetSelected(index int) {
	c.ResetSelected()
	c.selected = index
	c.connections[index].focused = true
	c.connections[index].Refresh()
	c.ChangeView(index)
	c.Refresh()
}

func (c *ConnectionList) ResetSelected() {
	c.selected = -1
	for _, obj := range c.connections {
		obj.focused = false
		obj.Refresh()
	}
	c.Refresh()
}

type connectionListRenderer struct {
	BaseRenderer

	background *canvas.Rectangle
	connections			[]*ConnectionButton
	layout     fyne.Layout
}

func (r *connectionListRenderer) MinSize() fyne.Size {
	l := len(r.connections)
	if l < 1 {
		return fyne.NewSize(0, 0)
	}

	h := float32(10)
	w := float32(10)

	mh := float32(0)
	for _, obj := range r.connections {
		s := obj.Size()
		if s.Height > mh {
			mh = s.Height
		}

		w = w + s.Width
	}
	h = h + mh
	w = w + float32((len(r.connections) - 1) * 5)

	return fyne.NewSize(w, h)
}

func (r *connectionListRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)

	x := float32(5)
	for _, obj := range r.connections {
		s := obj.MinSize()
		obj.Resize(s)
		obj.Move(fyne.NewPos(x, (size.Height - s.Height) / 2))
		x = x + s.Width + 5
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

