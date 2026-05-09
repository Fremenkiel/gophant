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

	Importance    widget.Importance

	OnTapped, OnDoubleTapped, OnTappedSecondary func(pe *fyne.PointEvent) `json:"-"`

	hovered, focused bool
	tapAnim          *fyne.Animation
	isAnimating      bool
}

func NewConnectionButton(handler *handlers.ConnectionHandler, l, d, r func(*fyne.PointEvent)) *ConnectionButton {
	t := canvas.NewText(handler.Connection.Name, nil)
	t.TextSize = 12

	i := canvas.NewCircle(th.Palette.Background)
	i.Resize(fyne.NewSize(12, 12))

	b := &ConnectionButton{label: t, icon: container.NewGridWrap(fyne.NewSize(12, 12), i), OnTapped: l, OnDoubleTapped: d, OnTappedSecondary: r}
	b.ExtendBaseWidget(b)
	return b
}

func (c *ConnectionButton) CreateRenderer() fyne.WidgetRenderer {
	c.ExtendBaseWidget(c)
	t := c.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	background := canvas.NewRectangle(t.Color(theme.ColorNameButton, v))
	background.CornerRadius = t.Size(theme.SizeNameInputRadius)

	tapBG := canvas.NewRectangle(color.Transparent)
	c.tapAnim = newButtonTapAnimation(tapBG, c, t)
	c.tapAnim.Curve = fyne.AnimationEaseOut

	objects := []fyne.CanvasObject{
		background,
		tapBG,
		c.label,
		c.icon,
	}
	b := &connectionButtonRenderer{
		BaseRenderer: NewBaseRenderer(objects),
		label: c.label,
		icon: c.icon,
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
	if b.OnTapped != nil {
		b.OnTapped(pe)
	}
}

func (b *ConnectionButton) DoubleTapped(pe *fyne.PointEvent) {
	if b.OnDoubleTapped != nil {
		b.OnDoubleTapped(pe)
	}
}

func (b *ConnectionButton) MouseIn(*desktop.MouseEvent) {
	b.hovered = true
	b.Refresh()
}

func (b *ConnectionButton) MouseOut() {
	b.hovered = false
	b.Refresh()
}

type connectionButtonRenderer struct {
	BaseRenderer

	background *canvas.Rectangle
	tapBG      *canvas.Rectangle
	label			*canvas.Text
	icon			*fyne.Container
	button     *ConnectionButton
	layout     fyne.Layout
}

func (r *connectionButtonRenderer) MinSize() fyne.Size {
	return r.label.MinSize()
}

func (r *connectionButtonRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	if !r.button.isAnimating {
		r.tapBG.Resize(size)
	}

	i := r.icon
	l := r.label

	is := fyne.NewSize(0, 0)
	if i != nil {
		is = i.MinSize()
		ip := fyne.NewPos(10, (size.Height - is.Height) / 2)
		i.Resize(is)
		i.Move(ip)
	}

	ls := l.MinSize()
	lp := fyne.NewPos(10 + is.Width + 5, (size.Height - ls.Height) / 2)
	l.Resize(ls)
	l.Move(lp)
}

func (r *connectionButtonRenderer) applyTheme() {
	t := r.button.Theme()
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
}

func (r *connectionButtonRenderer) Refresh() {
	r.applyTheme()
}

func (r *connectionButtonRenderer) buttonColorNames() (forground, background, backgroundBlend fyne.ThemeColorName) {
	b := r.button
	if b.focused {
		backgroundBlend = theme.ColorNameFocus
	} else if b.hovered {
		backgroundBlend = theme.ColorNameHover
	}

	return th.ColorNameButtonForeground, theme.ColorNameButton, backgroundBlend
}
