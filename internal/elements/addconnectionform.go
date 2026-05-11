package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/fs"
	"github.com/Fremenkiel/gophant/v2/internal/th"
)

type AddConnectionForm struct {
	widget.BaseWidget

	connections []*ConnectionButton

	selected int

	ChangeView	func(index int)
	AddConnection func(*fyne.PointEvent)
}

func NewAddConnectionForm(items []*ConnectionButton, changeView func(int), addConnection func(*fyne.PointEvent)) *AddConnectionForm {
	c := &AddConnectionForm{
		connections: items,
		selected: -1,
		ChangeView: changeView,
		AddConnection: addConnection,
	}
	c.ExtendBaseWidget(c)
	return c
}

func (c *AddConnectionForm) CreateRenderer() fyne.WidgetRenderer {
	c.ExtendBaseWidget(c)
	t := c.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	background := canvas.NewRectangle(t.Color(theme.ColorNameButton, v))
	background.CornerRadius = t.Size(theme.SizeNameInputRadius)

	header := canvas.NewText("CONNECTIONS", th.Palette.SecondaryText)
	header.TextSize = 10

	ab := NewIconButton(fs.IconNameAdd,
		func(pe *fyne.PointEvent) {
			c.AddConnection(pe)
		})

	l := widget.NewList(
		func() int {
			return len(c.connections)
		},
		func() fyne.CanvasObject {
			return NewConnectionButton("Template", "ro", nil, nil, nil)
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			b := c.connections[lii]
			b.OnTapped = func(pe *fyne.PointEvent) {
			c.SetSelected(lii)
			}
			co.(*ConnectionButton).SetContent(b.label.Text, b.pLabel.Text)
		},
		)
	l.HideSeparators = true

	objects := []fyne.CanvasObject{
		background,
		header,
		ab,
		l,
	}

	r := &addConnectionFormRenderer{
		BaseRenderer: NewBaseRenderer(objects),
		header: header,
		addButton: ab,
		cList: c,
		list: l,
		background: background,
	}
	r.applyTheme()
	return r
}

func (c *AddConnectionForm) SetContent(items	[]*ConnectionButton) {
	c.connections = items
	c.Refresh()
}

func (c *AddConnectionForm) SetSelected(index int) {
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

type addConnectionFormRenderer struct {
	BaseRenderer

	background *canvas.Rectangle
	header				*canvas.Text
	addButton				*IconButton
	cList			*AddConnectionForm
	list			*widget.List
	layout     fyne.Layout
}

func (r *addConnectionFormRenderer) MinSize() fyne.Size {
	hs := r.header.MinSize()

	h := hs.Height + 25
	l := len(r.cList.connections)
	if l < 1 {
		return fyne.NewSize(0, h)
	}

	h = h + float32(r.cList.connections[0].MinSize().Height * 4) + 1
	return fyne.NewSize(r.cList.connections[0].MinSize().Width, h)
}

func (r *addConnectionFormRenderer) Layout(size fyne.Size) {
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
	l.Resize(fyne.NewSize(size.Width - 2, size.Height - ch - 1))
	l.Move(fyne.NewPos(1, ch))
}

func (r *addConnectionFormRenderer) applyTheme() {
	t := fyne.CurrentApp().Settings().Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	if bg := r.background; bg != nil {
		bg.StrokeColor = t.Color(theme.ColorNameSeparator, v)
		bg.StrokeWidth = 1
		bg.CornerRadius = 0
	}
}

func (r *addConnectionFormRenderer) Refresh() {
	r.list.Refresh()
	r.applyTheme()
}

