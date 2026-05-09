package containers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/elements"
)

type SidebarTab struct {
	widget.BaseWidget

	tabs []*elements.SidebarTab

	selected, focused int

	ChangeView	func(index int)
}

func NewSidebarTab(changeView func(int), items ...*elements.SidebarTab) *SidebarTab {
	c := &SidebarTab{tabs: items, selected: -1, ChangeView: changeView}
	c.ExtendBaseWidget(c)
	return c
}

func (c *SidebarTab) CreateRenderer() fyne.WidgetRenderer {
	c.ExtendBaseWidget(c)
	t := c.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	background := canvas.NewRectangle(t.Color(theme.ColorNameButton, v))
	background.CornerRadius = t.Size(theme.SizeNameInputRadius)

	objects := []fyne.CanvasObject{
		background,
	}
	for i, obj := range c.tabs {
		if i == 0 {
			c.selected = 0
			obj.SetFocus(true)
		}
		obj.OnTapped = func(pe *fyne.PointEvent) {
			c.SetSelected(i)
		}
		objects = append(objects, obj)
	}

	r := &sidebarTabRenderer{
		BaseRenderer: elements.NewBaseRenderer(objects),
		tabs: c.tabs,
		background: background,
	}
	r.applyTheme()
	return r
}

func (c *SidebarTab) SetSelected(index int) {
	c.ResetSelected()
	c.selected = index
	c.tabs[index].SetFocus(true)
	c.tabs[index].Refresh()
	c.ChangeView(index)
	c.Refresh()
}

func (c *SidebarTab) ResetSelected() {
	c.selected = -1
	for _, obj := range c.tabs {
		obj.SetFocus(false)
		obj.Refresh()
	}
	c.Refresh()
}

type sidebarTabRenderer struct {
	elements.BaseRenderer

	background *canvas.Rectangle
	tabs			[]*elements.SidebarTab
	layout     fyne.Layout
}

func (r *sidebarTabRenderer) MinSize() fyne.Size {
	l := len(r.tabs)
	if l < 1 {
		return fyne.NewSize(0, 0)
	}

	h := float32(16)
	w := float32(10)

	mh := float32(0)
	for _, obj := range r.tabs {
		s := obj.Size()
		if s.Height > mh {
			mh = s.Height
		}

		w = w + s.Width
	}
	h = h + mh
	w = w + float32((len(r.tabs) - 1) * 5)

	return fyne.NewSize(w, h)
}

func (r *sidebarTabRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)

	x := float32(5)
	for _, obj := range r.tabs {
		s := obj.MinSize()
		obj.Resize(s)
		obj.Move(fyne.NewPos(x, (size.Height - s.Height) / 2))
		x = x + s.Width + 5
	}
}

func (r *sidebarTabRenderer) applyTheme() {
	t := fyne.CurrentApp().Settings().Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	if bg := r.background; bg != nil {
		bg.StrokeColor = t.Color(theme.ColorNameSeparator, v)
		bg.StrokeWidth = 1
		bg.CornerRadius = 0
	}
}

func (r *sidebarTabRenderer) Refresh() {
	r.applyTheme()
}

