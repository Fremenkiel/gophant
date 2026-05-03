package elements

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/utils"
)

type IconBox struct {
	widget.BaseWidget
	label			*widget.Label
	icon			*fyne.Container

	Importance    widget.Importance

		OnTapped, OnDoubleTapped, OnTappedSecondary func(pe *fyne.PointEvent) `json:"-"`

	hovered, focused bool
	tapAnim          *fyne.Animation
	isAnimating      bool
}

func NewIconBox(text string, icon fyne.CanvasObject, l, d, r func(*fyne.PointEvent)) *IconBox {
	if icon == nil {
		icon = canvas.NewCircle(color.Transparent)
	}
		icon.Resize(fyne.NewSize(12, 12))

	b := &IconBox{label: widget.NewLabel(text), icon: container.NewGridWrap(fyne.NewSize(12, 12), icon), OnTapped: l, OnDoubleTapped: d, OnTappedSecondary: r}
	b.ExtendBaseWidget(b)
	return b
}

func (i *IconBox) CreateRenderer() fyne.WidgetRenderer {
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
		i.icon,
	}
	b := &iconBoxRenderer{
		BaseRenderer: NewBaseRenderer(objects),
		label: i.label,
		icon: i.icon,
		button: i,
		background: background,
		tapBG: tapBG,
	}
	b.applyTheme()
	return b
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

func (b *IconBox) MouseIn(*desktop.MouseEvent) {
	b.hovered = true
	b.Refresh()
}

func (b *IconBox) MouseMoved(*desktop.MouseEvent) {
}

func (b *IconBox) MouseOut() {
	b.hovered = false
	b.Refresh()
}


type iconBoxRenderer struct {
	BaseRenderer

	background *canvas.Rectangle
	tapBG      *canvas.Rectangle
	label			*widget.Label
	icon			*fyne.Container
	button     *IconBox
	layout     fyne.Layout
}

func (r *iconBoxRenderer) MinSize() fyne.Size {
	return r.button.label.MinSize()
}

func (r *iconBoxRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	if !r.button.isAnimating {
		r.tapBG.Resize(size)
	}

	l := r.button.label
	ic := r.button.icon

	is := fyne.NewSize(0, 0)
	if ic != nil {
		is = ic.MinSize()
		ip := fyne.NewPos(5, (size.Height - is.Height) / 2)
		ic.Resize(is)
		ic.Move(ip)
	}

	ls := l.MinSize()
	lp := fyne.NewPos(10 + is.Width, 0)
	l.Resize(ls)
	l.Move(lp)
}

func (r *iconBoxRenderer) applyTheme() {
	th := r.button.Theme()
	_, bgColorName, bgBlendName := r.buttonColorNames()
	if bg := r.background; bg != nil {
		v := fyne.CurrentApp().Settings().ThemeVariant()
		bgColor := color.Color(color.Transparent)
		if bgColorName != "" {
			bgColor = th.Color(bgColorName, v)
		}
		if bgBlendName != "" {
			bgColor = blendColor(bgColor, th.Color(bgBlendName, v))
		}
		bg.FillColor = bgColor
		bg.CornerRadius = th.Size(theme.SizeNameInputRadius)
		bg.Refresh()
	}

	r.label.Refresh()
		r.icon.Refresh()
}

func (r *iconBoxRenderer) Refresh() {
	r.label.Refresh()
	if r.icon != nil {
		r.icon.Refresh()
	}
	r.applyTheme()
}

func (r *iconBoxRenderer) buttonColorNames() (foreground, background, backgroundBlend fyne.ThemeColorName) {
	foreground = theme.ColorNameForeground
	b := r.button
	if b.focused {
		backgroundBlend = theme.ColorNameFocus
	} else if b.hovered {
		backgroundBlend = theme.ColorNameHover
	}
	if background == "" {
		switch b.Importance {
		case widget.DangerImportance:
			foreground = theme.ColorNameForegroundOnError
			background = theme.ColorNameError
		case widget.HighImportance:
			foreground = theme.ColorNameForegroundOnPrimary
			background = theme.ColorNamePrimary
		case widget.LowImportance:
			if backgroundBlend != "" {
				background = theme.ColorNameButton
			}
		case widget.SuccessImportance:
			foreground = theme.ColorNameForegroundOnSuccess
			background = theme.ColorNameSuccess
		case widget.WarningImportance:
			foreground = theme.ColorNameForegroundOnWarning
			background = theme.ColorNameWarning
		default:
			background = theme.ColorNameButton
		}
	}
	return foreground, background, backgroundBlend
}

func blendColor(under, over color.Color) color.Color {
	dstR, dstG, dstB, dstA := under.RGBA()
	srcR, srcG, srcB, srcA := over.RGBA()

	srcAlpha := float32(srcA) / 0xFFFF
	dstAlpha := float32(dstA) / 0xFFFF

	outAlpha := srcAlpha + dstAlpha*(1-srcAlpha)
	outR := srcR + uint32(float32(dstR)*(1-srcAlpha))
	outG := srcG + uint32(float32(dstG)*(1-srcAlpha))
	outB := srcB + uint32(float32(dstB)*(1-srcAlpha))

	return color.RGBA64{R: uint16(outR), G: uint16(outG), B: uint16(outB), A: uint16(outAlpha * 0xFFFF)}
}

func newButtonTapAnimation(bg *canvas.Rectangle, w fyne.Widget, th fyne.Theme) *fyne.Animation {
	v := fyne.CurrentApp().Settings().ThemeVariant()
	return fyne.NewAnimation(canvas.DurationStandard, func(done float32) {
		mid := w.Size().Width / 2
		size := mid * done
		bg.Resize(fyne.NewSize(size*2, w.Size().Height))
		bg.Move(fyne.NewPos(mid-size, 0))

		r, g, bb, a := utils.ToNRGBA(th.Color(theme.ColorNamePressed, v))
		aa := uint8(a)
		fade := aa - uint8(float32(aa)*done)
		if fade > 0 {
			bg.FillColor = &color.NRGBA{R: uint8(r), G: uint8(g), B: uint8(bb), A: fade}
		} else {
			bg.FillColor = color.Transparent
		}
		canvas.Refresh(bg)
		if done == 1.0 {
			if btn, ok := w.(*IconBox); ok {
				btn.isAnimating = false
			}
		}
	})
}
