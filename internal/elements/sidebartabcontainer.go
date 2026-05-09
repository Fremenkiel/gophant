package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type SidebarTabContainer struct {
	widget.BaseWidget

	tabs []*SidebarTab

	selected int

	ChangeView	func(index int)
}

func NewSidebarTabContainer(changeView func(int), items ...*SidebarTab) *SidebarTabContainer {
	c := &SidebarTabContainer{tabs: items, selected: -1, ChangeView: changeView}
	c.ExtendBaseWidget(c)
	return c
}

func (c *SidebarTabContainer) CreateRenderer() fyne.WidgetRenderer {
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
			obj.focused = true
		}
		obj.OnTapped = func(pe *fyne.PointEvent) {
			c.SetSelected(i)
		}
		objects = append(objects, obj)
	}

	r := &sidebarTabContainerRenderer{
		BaseRenderer: NewBaseRenderer(objects),
		tabs: c.tabs,
		background: background,
	}
	r.applyTheme()
	return r
}

func (c *SidebarTabContainer) SetSelected(index int) {
	c.ResetSelected()
	c.selected = index
	c.tabs[index].focused = true
	c.tabs[index].Refresh()
	c.ChangeView(index)
	c.Refresh()
}

func (c *SidebarTabContainer) ResetSelected() {
	c.selected = -1
	for _, obj := range c.tabs {
		obj.focused = false
		obj.Refresh()
	}
	c.Refresh()
}

type sidebarTabContainerRenderer struct {
	BaseRenderer

	background *canvas.Rectangle
	tabs			[]*SidebarTab
	layout     fyne.Layout
}

func (r *sidebarTabContainerRenderer) MinSize() fyne.Size {
	l := len(r.tabs)
	if l < 1 {
		return fyne.NewSize(0, 0)
	}

	h := float32(10)
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

func (r *sidebarTabContainerRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)

	x := float32(5)
	for _, obj := range r.tabs {
		s := obj.MinSize()
		obj.Resize(s)
		obj.Move(fyne.NewPos(x, (size.Height - s.Height) / 2))
		x = x + s.Width + 5
	}
}

func (r *sidebarTabContainerRenderer) applyTheme() {
	t := fyne.CurrentApp().Settings().Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	if bg := r.background; bg != nil {
		bg.StrokeColor = t.Color(theme.ColorNameSeparator, v)
		bg.StrokeWidth = 1
		bg.CornerRadius = 0
	}

	/*
	t := r.Theme()
	b := r.button
	fgColorName, bgColorName, bgBlendName := r.buttonColorNames()

	if bg := r.background; bg != nil {
		v := fyne.CurrentApp().Settings().ThemeVariant()
		bgColor := color.Color(color.Transparent)
		if bgColorName != "" {
			bgColor = t.Color(bgColorName, v)
		}
		if bgBlendName != "" {
			bgColor = blendColor(bgColor, t.Color(bgBlendName, v))
		}
		bg.FillColor = bgColor
		bg.CornerRadius = t.Size(theme.SizeNameInputRadius)
		bg.Refresh()
	}

	if l := r.label; l != nil {
		v := fyne.CurrentApp().Settings().ThemeVariant()
		fgColor := color.Color(color.Transparent)
		if fgColorName != "" {
			fgColor = t.Color(fgColorName, v)
		}
		if b.focused || b.hovered {
			fgColor = t.Color(th.ColorNameFocusText, v)
		}
		l.Color = fgColor
		l.Refresh()
	}

	if i := r.icon; i != nil {
		fgColor := fgColorName
		if b.focused || b.hovered {
			fgColor = th.ColorNameFocusText
		}
		i.Objects[0] = r.button.createIcon(fgColor)
		i.Refresh()
	}
	*/
}

func (r *sidebarTabContainerRenderer) Refresh() {
	r.applyTheme()
}

