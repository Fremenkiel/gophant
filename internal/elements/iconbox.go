package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/themes"
)

type IconBox struct {
	widget.BaseWidget
	label			*widget.Label
	icon			*fyne.Container
	OnTapped, OnDoubleTapped, OnTappedSecondary func(pe *fyne.PointEvent)
}

func NewIconBox(text string, icon fyne.CanvasObject, l, d, r func(*fyne.PointEvent)) *IconBox {
	if icon == nil {
		icon = canvas.NewCircle(themes.Palette.Background)
	}
		icon.Resize(fyne.NewSize(12, 12))

	b := &IconBox{label: widget.NewLabel(text), icon: container.NewGridWrap(fyne.NewSize(12, 12), icon), OnTapped: l, OnDoubleTapped: d, OnTappedSecondary: r}
	b.ExtendBaseWidget(b)
	return b
}

func (i *IconBox) CreateRenderer() fyne.WidgetRenderer {
	var c *fyne.Container
		c = container.New(&themes.IconBox{}, i.label, i.icon)
	return widget.NewSimpleRenderer(c)
}

func (i *IconBox) SetContent(text string, icon fyne.CanvasObject) {
	i.label.SetText(text)

	if icon != nil {
		icon.Resize(fyne.NewSize(12, 12))
			i.icon.Objects[0] = icon
	}

	i.Refresh()
}

func (i *IconBox) TappedSecondary(pe *fyne.PointEvent) {
	if i.OnTappedSecondary != nil {
		i.OnTappedSecondary(pe)
	}
}

func (i *IconBox) Tapped(pe *fyne.PointEvent) {
	if i.OnTapped != nil {
		i.OnTapped(pe)
	}
}

func (i *IconBox) DoubleTapped(pe *fyne.PointEvent) {
	if i.OnDoubleTapped != nil {
		i.OnDoubleTapped(pe)
	}
}

func (i *IconBox) Refresh() {
	i.label.Refresh()
	if i.icon != nil {
		i.icon.Refresh()
	}
}
