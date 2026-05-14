package elements

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type InputAdornment struct {
	widget.BaseWidget
	Entry *focusEntry

	StartAdornment	fyne.CanvasObject
	EndAdornment		fyne.CanvasObject

	focus, hovered						bool
	StartSpacer, EndSpacer			bool

	override *container.ThemeOverride
}

func NewInputAdornment(startAdornment fyne.CanvasObject, endAdornment fyne.CanvasObject) *InputAdornment {
	i := &InputAdornment{
		StartAdornment: startAdornment,
		EndAdornment: endAdornment,
	}

	fe := &focusEntry{parent: i}
  fe.ExtendBaseWidget(fe)
  i.Entry = fe
  i.override = container.NewThemeOverride(fe, &inputAdornmentTheme{
		Theme: fyne.CurrentApp().Settings().Theme(),
	})
  i.ExtendBaseWidget(i)
	return i
}

func (i *InputAdornment) CreateRenderer() fyne.WidgetRenderer {
	i.ExtendBaseWidget(i)
	th := i.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	background := canvas.NewRectangle(th.Color(theme.ColorNameButton, v))
	background.CornerRadius = th.Size(theme.SizeNameInputRadius)

	ss := canvas.NewLine(color.Transparent)
	ss.StrokeWidth = 1
	ss.Hide()
	
	es := canvas.NewLine(color.Transparent)
	es.StrokeWidth = 1
	es.Hide()

	objects := []fyne.CanvasObject{
		background,
		i.Entry,
	}

	if i.StartAdornment != nil {
		objects = append(objects, i.StartAdornment)
	}

	if i.EndAdornment != nil {
		objects = append(objects, i.EndAdornment)
	}

	objects = append(objects, ss)
	objects = append(objects, es)

	b := &inputAdornmentRenderer{
		BaseRenderer: NewBaseRenderer(objects),
		background: background,
		input: i,
		startSpacer: ss,
		endSpacer: es,
	}
	b.applyTheme()
	return b
}

func (i *InputAdornment) SetText(s string)     { i.Entry.SetText(s) }
func (i *InputAdornment) Text() string         { return i.Entry.Text }
func (i *InputAdornment) SetPassword(b bool)   { i.Entry.Password = b; i.Entry.Refresh() }
func (i *InputAdornment) SetValidator(v fyne.StringValidator) { i.Entry.Validator = v }
func (i *InputAdornment) SetOnChanged(f func(string)) { i.Entry.OnChanged = f }

type inputAdornmentRenderer struct {
	BaseRenderer

	background *canvas.Rectangle
	input      *InputAdornment
	startSpacer	*canvas.Line
	endSpacer		*canvas.Line
	layout     fyne.Layout
}

func (r *inputAdornmentRenderer) MinSize() fyne.Size {
	s := r.input.Entry.MinSize()
	return fyne.NewSize(s.Width, s.Height + 2)
}

func (r *inputAdornmentRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)

	sp := float32(0)
	ep := float32(0)

	r.endSpacer.Resize(r.endSpacer.MinSize())

	sas := fyne.NewSize(0, 0)
	if r.input.StartAdornment != nil {
		sas = r.input.StartAdornment.MinSize()
		r.input.StartAdornment.Resize(sas)
		r.input.StartAdornment.Move(fyne.NewPos(10, (size.Height - sas.Height) / 2))
		if r.input.StartSpacer {
			r.startSpacer.Position1 = fyne.NewPos(0, 0)
			r.startSpacer.Position2 = fyne.NewPos(0, size.Height / 2)
			r.startSpacer.Move(fyne.NewPos(sas.Width + 20, size.Height / 2 / 2))
			sp = sas.Width + 31
		} else {
			sp = sas.Width + 21
		}
	}

	eas := fyne.NewSize(0, 0)
	if r.input.EndAdornment != nil {
		eas = r.input.EndAdornment.MinSize()
		r.input.EndAdornment.Resize(eas)
		r.input.EndAdornment.Move(fyne.NewPos(size.Width - eas.Width - 10, (size.Height - eas.Height) / 2))
		if r.input.EndSpacer {
			r.endSpacer.Position1 = fyne.NewPos(0, 0)
			r.endSpacer.Position2 = fyne.NewPos(0, size.Height / 2)
			r.endSpacer.Move(fyne.NewPos(size.Width - eas.Width - 20, size.Height / 2 / 2))
			ep = eas.Width + 31
		} else {
			ep = eas.Width + 21
		}
	}

	r.input.Entry.Resize(fyne.NewSize(size.Width - sp - ep, size.Height - 2))
	r.input.Entry.Move(fyne.NewPos(sp, 1))
}

func (r *inputAdornmentRenderer) applyTheme() {
	t := fyne.CurrentApp().Settings().Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	if bg := r.background; bg != nil {
		bg.FillColor = t.Color(theme.ColorNameInputBackground, v)
		bg.CornerRadius = t.Size(theme.SizeNameInputRadius)
		bg.StrokeWidth = 1
		if r.input.focus {
			bg.StrokeColor = t.Color(theme.ColorNamePrimary, v)
		} else {
			bg.StrokeColor = t.Color(theme.ColorNameInputBorder, v)
		}
		bg.Refresh()
	}

	if ss := r.startSpacer; ss != nil {
		ss.StrokeColor = t.Color(theme.ColorNameInputBorder, v)
		if r.input.StartSpacer {
			ss.Show()
		}
	}

	if es := r.endSpacer; es != nil {
		es.StrokeColor = t.Color(theme.ColorNameInputBorder, v)
		if r.input.EndSpacer {
			es.Show()
		}
	}
}

func (r *inputAdornmentRenderer) Refresh() {
	r.applyTheme()
}

type inputAdornmentTheme struct{ fyne.Theme }

func (t *inputAdornmentTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameInputBorder:        // unfocused border
		return color.Transparent
	case theme.ColorNamePrimary:            // focused border (Entry swaps to this)
		return color.Transparent
	case theme.ColorNameDisabled:           // disabled border
		return color.Transparent
	case theme.ColorNameForeground:
		if v == theme.VariantLight {
			return t.Theme.Color(n, v)
		}
		return color.RGBA{229, 229, 233, 255}
	case theme.ColorNameSelection:
		return color.RGBA{255, 10, 10, 255}
	}
	return t.Theme.Color(n, v)
}

func (t *inputAdornmentTheme) Size(n fyne.ThemeSizeName) float32 {
	switch n {
	case theme.SizeNameInputBorder:   // border width (0 = no border)
		return 1
	case theme.SizeNameInputRadius:   // corner radius
		return 0
	case theme.SizeNameInnerPadding:  // text inset from border
		return 6
	case theme.SizeNameText:
		return 12
	}
	return t.Theme.Size(n)
}

type focusEntry struct {
	widget.Entry
	parent *InputAdornment
}

func (e *focusEntry) FocusGained() {
	e.Entry.FocusGained()
	e.parent.focus = true
	e.parent.Refresh()
}

func (e *focusEntry) FocusLost() {
	e.Entry.FocusLost()
	e.parent.focus = false
	e.parent.Refresh()
}

