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

type IconButton struct {
	widget.BaseWidget
	icon			*fyne.StaticResource

	Importance    widget.Importance

	OnTapped func(pe *fyne.PointEvent) `json:"-"`

	hovered, focused bool
	tapAnim          *fyne.Animation
	isAnimating      bool
}

func NewIconButton(icon string, l func(*fyne.PointEvent)) *IconButton {
	c, err := fs.StaticFS.ReadFile(fmt.Sprint("static/", icon))
	if err != nil {
		log.Println("Unable to read file")
		log.Fatal(err)
	}
	res := fyne.NewStaticResource(icon, c)

	b := &IconButton{icon: res, OnTapped: l}
	b.ExtendBaseWidget(b)
	return b
}

func (i *IconButton) CreateRenderer() fyne.WidgetRenderer {
	i.ExtendBaseWidget(i)
	t := i.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	background := canvas.NewRectangle(t.Color(theme.ColorNameButton, v))

	tapBG := canvas.NewRectangle(color.Transparent)
	i.tapAnim = newButtonTapAnimation(tapBG, i, t)
	i.tapAnim.Curve = fyne.AnimationEaseOut

	icon := container.NewGridWrap(fyne.NewSize(14, 14), i.createIcon(th.ColorNameButtonForeground))

	objects := []fyne.CanvasObject{
		background,
		tapBG,
		icon,
	}
	b := &iconButtonRenderer{
		BaseRenderer: NewBaseRenderer(objects),
		icon: icon,
		button: i,
		background: background,
		tapBG: tapBG,
	}
	b.applyTheme()
	return b
}

func (b *IconButton) TappedSecondary(pe *fyne.PointEvent) {
}

func (b *IconButton) Tapped(pe *fyne.PointEvent) {
	if b.OnTapped != nil {
		b.OnTapped(pe)
	}
}

func (b *IconButton) MouseIn(me *desktop.MouseEvent) {
	b.hovered = true
	b.Refresh()
}

func (b *IconButton) MouseMoved(me *desktop.MouseEvent) {
}

func (b *IconButton) MouseOut() {
	b.hovered = false
	b.Refresh()
}

func (b *IconButton) createIcon(c fyne.ThemeColorName) *canvas.Image {
	res := b.icon
	s := fyne.NewSize(14, 14)

	t := th.NewColoredResource(res, th.ColorNameTransparent, c)
	i := canvas.NewImageFromResource(t)
	i.Resize(s)
	return i
}

type iconButtonRenderer struct {
	BaseRenderer

	background *canvas.Rectangle
	tapBG      *canvas.Rectangle
	icon			*fyne.Container
	button     *IconButton
	layout     fyne.Layout
}

func (r *iconButtonRenderer) MinSize() fyne.Size {
	s := r.icon.MinSize()
	return fyne.NewSize(s.Width + 10, s.Height + 10)
}

func (r *iconButtonRenderer) Layout(size fyne.Size) {
	s := r.MinSize()

	r.background.Resize(s)
	if !r.button.isAnimating {
		r.tapBG.Resize(s)
	}

	ic := r.icon

	is := ic.MinSize()
	ip := fyne.NewPos((s.Width - is.Width) / 2, (s.Height - is.Height) / 2)
	ic.Move(ip)
}
func (r *iconButtonRenderer) applyTheme() {
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

	if i := r.icon; i != nil {
		fgColor := fgColorName
		if b.focused || b.hovered {
			fgColor = th.ColorNameFocusText
		}
		i.Objects[0] = r.button.createIcon(fgColor)
		i.Refresh()
	}
}

func (r *iconButtonRenderer) Refresh() {
	r.applyTheme()
}

func (r *iconButtonRenderer) buttonColorNames() (forground, background, backgroundBlend fyne.ThemeColorName) {
	b := r.button
	if b.focused {
		backgroundBlend = theme.ColorNameFocus
	} else if b.hovered {
		backgroundBlend = theme.ColorNameHover
	}

	return th.ColorNameButtonForeground, theme.ColorNameButton, backgroundBlend
}

