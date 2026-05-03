package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Collapse struct {
	widget.BaseWidget
	button		*IconBox
	list			*fyne.Container
	open			bool
}

func NewCollapse(button *IconBox, list []*IconBox) *Collapse {
	li := container.NewVBox(widget.NewLabel("Placeholder"))
	li.RemoveAll()
	for _, el := range list {
		li.Add(el)
	}
		li.Hide()

	b := &Collapse{button: button, list: li, open: false	}
	b.ExtendBaseWidget(b)
	return b
}

func (c *Collapse) CreateRenderer() fyne.WidgetRenderer {
	co := container.NewBorder(c.button, nil, nil, nil, c.list)
	return widget.NewSimpleRenderer(co)
}

func (c *Collapse) SetContent(list []*IconBox) {
	if c.open {
		c.list.Show()
	} else {
		c.list.Hide()
	}

	c.Refresh()
}

func (c *Collapse) SetHeader(text string, icon fyne.CanvasObject, l, d, r func(pe *fyne.PointEvent)) {
	c.button.SetContent(text, icon)
	c.button.OnTapped = l
	c.button.OnDoubleTapped = d
	c.button.OnTappedSecondary = r
	c.Refresh()
}

func (c *Collapse) Toggle() {
	if c.open { c.Close() } else { c.Open() }
}

func (c *Collapse) Open() {
	c.open = true
	c.list.Show()
	c.list.Refresh()
}

func (c *Collapse) Close() {
	c.open = false
	c.list.Hide()
	c.list.Refresh()
}

func (c *Collapse) Refresh() {
	c.button.Refresh()
	c.list.Refresh()
}
