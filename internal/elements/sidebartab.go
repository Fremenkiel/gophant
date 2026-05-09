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

type SidebarTab struct {
	widget.BaseWidget
	label			*canvas.Text
	icon			*fyne.StaticResource

	Importance    widget.Importance

	OnTapped func(pe *fyne.PointEvent) `json:"-"`

	hovered, focused bool
	tapAnim          *fyne.Animation
	isAnimating      bool
}

func NewSidebarTab(text, icon string, l func(*fyne.PointEvent)) *SidebarTab {
	c, err := fs.StaticFS.ReadFile(fmt.Sprint("static/", icon))
	if err != nil {
		log.Println("Unable to read file")
		log.Fatal(err)
	}
	res := fyne.NewStaticResource(icon, c)

	t := canvas.NewText(text, nil)
	t.TextSize = 12

	b := &SidebarTab{label: t, icon: res}
	b.ExtendBaseWidget(b)
	return b
}

func (i *SidebarTab) CreateRenderer() fyne.WidgetRenderer {
	i.ExtendBaseWidget(i)
	t := i.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	background := canvas.NewRectangle(t.Color(theme.ColorNameButton, v))
	background.CornerRadius = t.Size(theme.SizeNameInputRadius)

	tapBG := canvas.NewRectangle(color.Transparent)
	i.tapAnim = newButtonTapAnimation(tapBG, i, t)
	i.tapAnim.Curve = fyne.AnimationEaseOut

	icon := container.NewGridWrap(fyne.NewSize(14, 14), i.createIcon(th.ColorNameButtonForeground))

	objects := []fyne.CanvasObject{
		background,
		tapBG,
		i.label,
		icon,
	}
	b := &sidebarTabRenderer{
		BaseRenderer: NewBaseRenderer(objects),
		label: i.label,
		icon: icon,
		button: i,
		background: background,
		tapBG: tapBG,
	}
	b.applyTheme()
	return b
}

func (i *SidebarTab) SetContent(text, icon string) {
	c, err := fs.StaticFS.ReadFile(fmt.Sprint("static/", icon))
	if err != nil {
		log.Println("Unable to read file")
		log.Fatal(err)
	}

	i.label.Text = text
	i.icon = fyne.NewStaticResource(icon, c)

	i.Refresh()
}

func (i *SidebarTab) Tapped(pe *fyne.PointEvent) {
	if i.OnTapped != nil {
		i.OnTapped(pe)
	}
}
func (b *SidebarTab) MouseIn(*desktop.MouseEvent) {
	b.hovered = true
	b.Refresh()
}

func (b *SidebarTab) MouseMoved(*desktop.MouseEvent) {
}

func (b *SidebarTab) MouseOut() {
	b.hovered = false
	b.Refresh()
}

func (b *SidebarTab) createIcon(c fyne.ThemeColorName) *canvas.Image {
	res := b.icon
	s := fyne.NewSize(14, 14)

	t := th.NewColoredResource(res, th.ColorNameTransparent, c)
	i := canvas.NewImageFromResource(t)
	i.Resize(s)
	return i
}


type sidebarTabRenderer struct {
	BaseRenderer

	background *canvas.Rectangle
	tapBG      *canvas.Rectangle
	label			*canvas.Text
	icon			*fyne.Container
	button     *SidebarTab
	layout     fyne.Layout
}

func (r *sidebarTabRenderer) MinSize() fyne.Size {
	h := r.icon.MinSize().Height + 15
	w := r.button.label.MinSize().Width + r.icon.MinSize().Width + 25
	return fyne.NewSize(w, h)
}

func (r *sidebarTabRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	if !r.button.isAnimating {
		r.tapBG.Resize(size)
	}

	l := r.button.label
	ic := r.icon

	is := fyne.NewSize(0, 0)
	if ic != nil {
		is = ic.MinSize()
		ip := fyne.NewPos(10, (size.Height - is.Height) / 2)
		ic.Resize(is)
		ic.Move(ip)
	}

	ls := l.MinSize()
	lp := fyne.NewPos(10 + is.Width + 5, (size.Height - ls.Height) / 2)
	l.Resize(ls)
	l.Move(lp)
}

func (r *sidebarTabRenderer) applyTheme() {
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

	if i := r.icon; i != nil {
		fgColor := fgColorName
		if b.focused || b.hovered {
			fgColor = th.ColorNameFocusText
		}
		i.Objects[0] = r.button.createIcon(fgColor)
		i.Refresh()
	}
}

func (r *sidebarTabRenderer) Refresh() {
	r.applyTheme()
}

func (r *sidebarTabRenderer) buttonColorNames() (forground, background, backgroundBlend fyne.ThemeColorName) {
	b := r.button
	if b.focused {
		backgroundBlend = theme.ColorNameFocus
	} else if b.hovered {
		backgroundBlend = theme.ColorNameHover
	}

	return th.ColorNameButtonForeground, theme.ColorNameButton, backgroundBlend
}
