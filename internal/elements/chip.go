package elements

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/th"
)

type Chip struct {
	widget.BaseWidget
	label			*canvas.Text
	icon			*fyne.Container

	Importance    widget.Importance

	OnTapped, OnTappedSecondary func(pe *fyne.PointEvent) `json:"-"`

	hovered, focused bool
	tapAnim          *fyne.Animation
	isAnimating      bool
}

func NewChip(text string, co color.Color, l, r func(*fyne.PointEvent)) *Chip {
	var icon *fyne.Container
	if co != nil {
		c := canvas.NewCircle(co)
		icon = container.NewGridWrap(fyne.NewSize(6, 6), c)
	}

	t := canvas.NewText(text, color.Transparent)
	t.TextSize = 11

	b := &Chip{label: t, icon: icon, OnTapped: l,  OnTappedSecondary: r}
	b.ExtendBaseWidget(b)
	return b
}

func (i *Chip) CreateRenderer() fyne.WidgetRenderer {
	i.ExtendBaseWidget(i)
	th := i.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	background := canvas.NewRectangle(th.Color(theme.ColorNameButton, v))
	background.CornerRadius = th.Size(theme.SizeNameInputRadius)

	tapBG := canvas.NewRectangle(color.Transparent)
	i.tapAnim = newButtonTapAnimation(tapBG, i, th)
	i.tapAnim.Curve = fyne.AnimationEaseOut

	objects := []fyne.CanvasObject{
		background,
		tapBG,
		i.label,
	}

	if i.icon != nil {
		objects = append(objects, i.icon)
	}

	b := &chipRenderer{
		BaseRenderer: NewBaseRenderer(objects),
		chip: i,
		background: background,
		tapBG: tapBG,
	}
	b.applyTheme()
	return b
}

func (i *Chip) TappedSecondary(pe *fyne.PointEvent) {
	if i.OnTappedSecondary != nil {
		i.OnTappedSecondary(pe)
	}
}

func (i *Chip) Tapped(pe *fyne.PointEvent) {
	if i.OnTapped != nil {
		i.OnTapped(pe)
	}
}

func (b *Chip) MouseIn(*desktop.MouseEvent) {
	b.hovered = true
	b.Refresh()
}

func (b *Chip) MouseMoved(*desktop.MouseEvent) {
}

func (b *Chip) MouseOut() {
	b.hovered = false
	b.Refresh()
}

func (c *Chip) SetFocus(state bool) {
	c.focused = state
}

type chipRenderer struct {
	BaseRenderer

	background *canvas.Rectangle
	tapBG      *canvas.Rectangle
	chip     *Chip
	layout     fyne.Layout
}

func (r *chipRenderer) MinSize() fyne.Size {
	ls := r.chip.label.MinSize()
	is := fyne.NewSize(0, 0)
	if r.chip.icon != nil {
		is = r.chip.icon.MinSize()
	}

	return fyne.NewSize(ls.Width + is.Width + 38, ls.Height + 6)
}

func (r *chipRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	if !r.chip.isAnimating {
		r.tapBG.Resize(size)
	}

	is := fyne.NewSize(0, 0)
	if ic := r.chip.icon; ic != nil {
		is = ic.MinSize()
		ip := fyne.NewPos(15, (size.Height - is.Height) / 2)
		ic.Resize(is)
		ic.Move(ip)
	}

	l := r.chip.label
	ls := l.MinSize()
	lp := fyne.NewPos(23 + is.Width, (size.Height - ls.Height) / 2)
	l.Resize(ls)
	l.Move(lp)
}

func (r *chipRenderer) applyTheme() {
	v := fyne.CurrentApp().Settings().ThemeVariant()
	t := r.chip.Theme()

	if bg := r.background; bg != nil {
		bg.FillColor = t.Color(theme.ColorNameInputBackground, v)
		bg.CornerRadius = 50
		bg.StrokeWidth = 1
		if r.chip.focused {
			bg.StrokeColor = t.Color(theme.ColorNamePrimary, v)
		} else if r.chip.hovered {
			bg.StrokeColor = t.Color(th.ColorNameInputBorderHover, v)
		} else {
			bg.StrokeColor = t.Color(theme.ColorNameInputBorder, v)
		}
		bg.Refresh()
	}

	if tbg := r.tapBG; tbg != nil {
		tbg.CornerRadius = 50
		tbg.StrokeWidth = 1
		if r.chip.focused {
			co := t.Color(theme.ColorNamePrimary, v)
			if icon := r.chip.icon; icon != nil {
				co = icon.Objects[0].(*canvas.Circle).FillColor
			}
			r, g, b, _ := co.RGBA()
			tbg.FillColor = color.RGBA{uint8(r), uint8(g), uint8(b), 32}
			tbg.StrokeColor = color.RGBA{uint8(r), uint8(g), uint8(b), 128}
		} else {
			tbg.FillColor = color.Transparent
			tbg.StrokeColor = color.Transparent
		}
		tbg.Refresh()
	}

	if label := r.chip.label; label != nil {
		if r.chip.focused || r.chip.hovered {
			label.Color = t.Color(th.ColorNameFocusText, v)
		} else {
			label.Color = t.Color(th.ColorNameText, v)
		}
		label.Refresh()
	}

	if icon := r.chip.icon; icon != nil {
		icon.Refresh()
	}
}

func (r *chipRenderer) Refresh() {
	r.applyTheme()
}
