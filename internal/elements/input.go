package elements

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Input struct {
	widget.BaseWidget
	Entry *widget.Entry
	override *container.ThemeOverride
}

func NewInput() *Input {
	i := &Input{Entry: widget.NewEntry()}
	i.override = container.NewThemeOverride(i.Entry, &inputTheme{
		Theme: fyne.CurrentApp().Settings().Theme(),
	})
	i.ExtendBaseWidget(i)
	return i
}

func (i *Input) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(i.override)
}

func (i *Input) SetText(s string)     { i.Entry.SetText(s) }
func (i *Input) Text() string         { return i.Entry.Text }
func (i *Input) SetPassword(b bool)   { i.Entry.Password = b; i.Entry.Refresh() }
func (i *Input) SetValidator(v fyne.StringValidator) { i.Entry.Validator = v }
func (i *Input) SetOnChanged(f func(string)) { i.Entry.OnChanged = f }

type inputTheme struct{ fyne.Theme }

func (t *inputTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameDisabled:           // disabled border
		if v == theme.VariantLight {
			return t.Theme.Color(n, v)
		}
		return color.RGBA{80, 80, 80, 255}
	case theme.ColorNameForeground:
		if v == theme.VariantLight {
			return t.Theme.Color(n, v)
		}
		return color.RGBA{229, 229, 233, 255}
	}
	return t.Theme.Color(n, v)
}

func (t *inputTheme) Size(n fyne.ThemeSizeName) float32 {
	switch n {
	case theme.SizeNameInputBorder:   // border width (0 = no border)
		return 1
	case theme.SizeNameInputRadius:   // corner radius
		return 4
	case theme.SizeNameInnerPadding:  // text inset from border
		return 6
	case theme.SizeNameText:
		return 12
	}
	return t.Theme.Size(n)
}
