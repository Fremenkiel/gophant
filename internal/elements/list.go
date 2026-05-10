package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type List struct {
	widget.BaseWidget

	Length	func() int
	RenderItem	func(lii widget.ListItemID, co fyne.CanvasObject)
}

func NewList(length func() int, renderItem func(lii widget.ListItemID, co fyne.CanvasObject)) *List {
	c := &List{
		Length: length,
		RenderItem: renderItem,
	}
	c.ExtendBaseWidget(c)
	return c
}

func (l *List) CreateRenderer() fyne.WidgetRenderer {
	l.ExtendBaseWidget(l)
	t := l.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	background := canvas.NewRectangle(t.Color(theme.ColorNameButton, v))
	background.CornerRadius = t.Size(theme.SizeNameInputRadius)

	objects := []fyne.CanvasObject{
		background,
	}

	r := &listRenderer{
		BaseRenderer: NewBaseRenderer(objects),
		list: l,
		background: background,
	}
	r.applyTheme()
	return r
}

func (c *List) SetContent(items	[]*ConnectionButton) {
	c.connections = items
	c.Refresh()
}

func (c *List) SetSelected(index int) {
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

type listRenderer struct {
	BaseRenderer

	background *canvas.Rectangle
	list			*List
	layout     fyne.Layout
}

func (r *listRenderer) MinSize() fyne.Size {
	hs := r.header.MinSize()

	h := hs.Height + 25
	l := len(r.cList.connections)
	if l < 1 {
		return fyne.NewSize(0, h)
	}

	h = h + float32(r.cList.connections[0].MinSize().Height * 4) + 1
	return fyne.NewSize(0, h)
}

func (r *listRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)

	h := r.header
	hs := fyne.NewSize(size.Width, h.MinSize().Height + 20)
	hp := fyne.NewPos(10, (hs.Height - h.MinSize().Height) / 2)
	h.Move(hp)

	ab := r.addButton
	ab.Resize(ab.MinSize())
	abp := fyne.NewPos(size.Width - ab.MinSize().Width - 5, (hs.Height - ab.MinSize().Height) / 2)
	ab.Move(abp)

	ch := float32(hs.Height + 5)
	l := r.list
	l.Resize(fyne.NewSize(size.Width - 2, size.Height - ch))
	l.Move(fyne.NewPos(1, ch))
}

func (r *listRenderer) applyTheme() {
	t := fyne.CurrentApp().Settings().Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	if bg := r.background; bg != nil {
		bg.StrokeColor = t.Color(theme.ColorNameSeparator, v)
		bg.StrokeWidth = 1
		bg.CornerRadius = 0
	}
}

func (r *listRenderer) Refresh() {
	/*
	objects := []fyne.CanvasObject{
		r.background,
		r.header,
		r.addButton,
	}
	for i, obj := range r.list.connections {
		obj.OnTapped = func(pe *fyne.PointEvent) {
			r.list.SetSelected(i)
		}
		objects = append(objects, obj)
	}
	r.SetObjects(objects)
	*/

	r.list.Refresh()
	r.applyTheme()
	//r.Layout(r.MinSize())
}

