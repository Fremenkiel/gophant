package elements

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/handlers"
	"github.com/Fremenkiel/gophant/v2/internal/th"
)

type ConnectionButton struct {
	widget.BaseWidget
	label			*canvas.Text
	icon			*fyne.Container
	pLabel			*canvas.Text

	Importance    widget.Importance

	OnTapped, OnDoubleTapped, OnTappedSecondary func(pe *fyne.PointEvent) `json:"-"`

	hovered, focused bool
	tapAnim          *fyne.Animation
	isAnimating      bool
}

func NewConnectionButton(handler *handlers.ConnectionHandler, l, d, r func(*fyne.PointEvent)) *ConnectionButton {
	t := canvas.NewText(handler.Connection.Name, nil)
	t.TextSize = 12

	i := canvas.NewCircle(th.Palette.Disabled)
	i.Resize(fyne.NewSize(6, 6))

	pt := canvas.NewText(handler.Connection.Permission, nil)
	pt.TextSize = 10

	b := &ConnectionButton{label: t, icon: container.NewGridWrap(fyne.NewSize(6, 6), i), pLabel: pt, OnTapped: l, OnDoubleTapped: d, OnTappedSecondary: r}
	b.ExtendBaseWidget(b)
	return b
}

func (c *ConnectionButton) CreateRenderer() fyne.WidgetRenderer {
	c.ExtendBaseWidget(c)
	t := c.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	background := canvas.NewRectangle(t.Color(theme.ColorNameButton, v))

	tapBG := canvas.NewRectangle(color.Transparent)
	c.tapAnim = newButtonTapAnimation(tapBG, c, t)
	c.tapAnim.Curve = fyne.AnimationEaseOut

	in := canvas.NewRectangle(th.Palette.Indicator)
	in.TopRightCornerRadius = 4
	in.BottomRightCornerRadius = 4
	in.Hide()

	objects := []fyne.CanvasObject{
		background,
		tapBG,
		c.label,
		c.icon,
		c.pLabel,
		in,
	}
	b := &connectionButtonRenderer{
		BaseRenderer: NewBaseRenderer(objects),
		label: c.label,
		icon: c.icon,
		pLabel: c.pLabel,
		indicator: in,
		button: c,
		background: background,
		tapBG: tapBG,
	}
	b.applyTheme()
	return b
}

func (b *ConnectionButton) TappedSecondary(pe *fyne.PointEvent) {
	if b.OnTappedSecondary != nil {
		b.OnTappedSecondary(pe)
	}
}

func (b *ConnectionButton) Tapped(pe *fyne.PointEvent) {
}

func (b *ConnectionButton) DoubleTapped(pe *fyne.PointEvent) {
	if b.OnDoubleTapped != nil {
		b.OnDoubleTapped(pe)
	}
}

func (b *ConnectionButton) MouseIn(pe *desktop.MouseEvent) {
	b.hovered = true
	b.Refresh()
}

func (b *ConnectionButton) MouseMoved(pe *desktop.MouseEvent) {
}

func (b *ConnectionButton) MouseOut() {
	b.hovered = false
	b.Refresh()
}

func (b *ConnectionButton) MouseDown(pe *desktop.MouseEvent) {
	if b.OnTapped != nil {
		b.OnTapped(nil)
	}
}

func (b *ConnectionButton) MouseUp(pe *desktop.MouseEvent) {}

func (t *ConnectionButton) SetFocus(focus bool) {
	t.focused = focus;
}

type connectionButtonRenderer struct {
	BaseRenderer

	background *canvas.Rectangle
	tapBG      *canvas.Rectangle
	label			*canvas.Text
	icon			*fyne.Container
	pLabel			*canvas.Text
	indicator	*canvas.Rectangle
	button     *ConnectionButton
	layout     fyne.Layout
}

func (r *connectionButtonRenderer) MinSize() fyne.Size {
	is := r.icon.MinSize()
	return fyne.NewSize(0, is.Height + 20)
}

func (r *connectionButtonRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	if !r.button.isAnimating {
		r.tapBG.Resize(size)
	}

	i := r.icon
	l := r.label
	in := r.indicator
	pl := r.pLabel

	if in != nil {
		s := fyne.NewSize(2, size.Height - 10)
		in.Resize(s)
		in.Move(fyne.NewPos(0, (size.Height - s.Height) / 2))
	}

	is := fyne.NewSize(0, 0)
	p := fyne.NewPos(12, 0)
	if i != nil {
		is = i.MinSize()
		p := fyne.NewPos(12, (size.Height - is.Height) / 2)
		i.Resize(is)
		i.Move(p)
	}

	ls := l.MinSize()
	lp := fyne.NewPos(p.X + is.Width + 12, (size.Height - ls.Height) / 2)
	l.Resize(ls)
	l.Move(lp)

	pls := l.MinSize()
	plp := fyne.NewPos(size.Width - pls.Width - 12, (size.Height - ls.Height) / 2)
	pl.Resize(pls)
	pl.Move(plp)
}

func (r *connectionButtonRenderer) applyTheme() {
	t := r.button.Theme()
	b := r.button

	if bg := r.background; bg != nil {
		v := fyne.CurrentApp().Settings().ThemeVariant()
		bgColor := t.Color(theme.ColorNameButton, v)
		if b.focused || b.hovered {
			bgColor = t.Color(th.ColorNameButtonHover, v)
		}
		bg.FillColor = bgColor
		bg.Refresh()
	}

	if l := r.label; l != nil {
		v := fyne.CurrentApp().Settings().ThemeVariant()
		fgColor := t.Color(th.ColorNameButtonForeground, v)
		if b.focused || b.hovered {
			fgColor = t.Color(th.ColorNameFocusText, v)
		}
		l.Color = fgColor
		l.Refresh()
	}

	if i := r.indicator; i != nil {
		if b.focused {
			i.Show()
		}
		i.Refresh()
	}
}

func (r *connectionButtonRenderer) Refresh() {
	r.applyTheme()
}

