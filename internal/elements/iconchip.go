package elements

import (
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/fs"
	"github.com/Fremenkiel/gophant/v2/internal/th"
)

type IconChip struct {
	widget.BaseWidget
	label			*canvas.Text
	icon			*fyne.StaticResource

	Importance    widget.Importance

	OnTapped, OnTappedSecondary func(pe *fyne.PointEvent) `json:"-"`

	hovered, focused bool
	tapAnim          *fyne.Animation
	isAnimating      bool
}

func NewIconChip(text string, icon string, l, r func(*fyne.PointEvent)) *IconChip {
	c, err := fs.StaticFS.ReadFile(fmt.Sprint("static/", icon))
	if err != nil {
		log.Println("Unable to read file")
		log.Fatal(err)
	}
	res := fyne.NewStaticResource(icon, c)

	t := canvas.NewText(text, color.Transparent)
	t.TextSize = 11

	b := &IconChip{label: t, icon: res, OnTapped: l,  OnTappedSecondary: r}
	b.ExtendBaseWidget(b)
	return b
}

func (i *IconChip) CreateRenderer() fyne.WidgetRenderer {
	i.ExtendBaseWidget(i)
	t := i.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	background := canvas.NewRectangle(t.Color(theme.ColorNameButton, v))
	background.CornerRadius = t.Size(theme.SizeNameInputRadius)

	icon := container.NewGridWrap(fyne.NewSize(14, 14), i.createIcon(th.ColorNameButtonForeground))

	objects := []fyne.CanvasObject{
		background,
		i.label,
		icon,
	}

	b := &iconChipRenderer{
		BaseRenderer: NewBaseRenderer(objects),
		chip: i,
		icon: icon,
		background: background,
	}
	b.applyTheme()
	return b
}

func (i *IconChip) TappedSecondary(pe *fyne.PointEvent) {
	if i.OnTappedSecondary != nil {
		i.OnTappedSecondary(pe)
	}
}

func (i *IconChip) Tapped(pe *fyne.PointEvent) {
	if i.OnTapped != nil {
		i.OnTapped(pe)
	}
}

func (b *IconChip) MouseIn(*desktop.MouseEvent) {
	b.hovered = true
	b.Refresh()
}

func (b *IconChip) MouseMoved(*desktop.MouseEvent) {
}

func (b *IconChip) MouseOut() {
	b.hovered = false
	b.Refresh()
}

func (c *IconChip) SetFocus(state bool) {
	c.focused = state
}

func (b *IconChip) createIcon(c fyne.ThemeColorName) *canvas.Image {
	res := b.icon
	s := fyne.NewSize(14, 14)

	t := th.NewColoredResource(res, th.ColorNameTransparent, c)
	i := canvas.NewImageFromResource(t)
	i.Resize(s)
	return i
}

type iconChipRenderer struct {
	BaseRenderer

	background *canvas.Rectangle
	chip     *IconChip
	icon			*fyne.Container
	layout     fyne.Layout
}

func (r *iconChipRenderer) MinSize() fyne.Size {
	ls := r.chip.label.MinSize()

	return fyne.NewSize(ls.Width + r.icon.MinSize().Width + 35, ls.Height + 6)
}

func (r *iconChipRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)

	ic := r.icon
	is := ic.MinSize()
	ip := fyne.NewPos(15, (size.Height - is.Height) / 2)
	ic.Resize(is)
	ic.Move(ip)

	l := r.chip.label
	ls := l.MinSize()
	lp := fyne.NewPos(20 + is.Width, (size.Height - ls.Height) / 2)
	l.Resize(ls)
	l.Move(lp)
}

func (r *iconChipRenderer) applyTheme() {
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

	if label := r.chip.label; label != nil {
		if r.chip.focused || r.chip.hovered {
			label.Color = t.Color(th.ColorNameFocusText, v)
		} else {
			label.Color = t.Color(th.ColorNameText, v)
		}
		label.Refresh()
	}

	if icon := r.icon; icon != nil {
		icon.Refresh()
	}
}

func (r *iconChipRenderer) Refresh() {
	r.applyTheme()
}
