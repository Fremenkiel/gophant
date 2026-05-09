package th

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type GophantTheme struct {}

var _ fyne.Theme = (*GophantTheme)(nil)

const (
	ColorNameIcon = "ColorNameIcon"
	ColorNameText = "ColorNameText"
	ColorNameButtonForeground = "ColorNameButtonForeground"
	ColorNameFocusText = "ColorNameFocusText"
)

func (m *GophantTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		if variant == theme.VariantLight {
			return theme.DefaultTheme().Color(name, variant)
		}
		return color.RGBA{8, 8, 10, 255}
	case theme.ColorNameForeground:
		if variant == theme.VariantLight {
			return theme.DefaultTheme().Color(name, variant)
		}
		return color.RGBA{168, 168, 178, 255}
	case theme.ColorNameSeparator:
		if variant == theme.VariantLight {
			return theme.DefaultTheme().Color(name, variant)
		}
		return color.RGBA{225, 50, 50, 255}
		// return color.RGBA{26, 26, 33, 0}
	case theme.ColorNameButton:
		if variant == theme.VariantLight {
			return theme.DefaultTheme().Color(name, variant)
		}
		return color.RGBA{13, 13, 16, 255}
	case ColorNameIcon:
		if variant == theme.VariantLight {
			return theme.DefaultTheme().Color(name, variant)
		}
		return color.RGBA{168, 168, 178, 255}
	case ColorNameText:
		if variant == theme.VariantLight {
			return theme.DefaultTheme().Color(name, variant)
		}
		return color.RGBA{168, 168, 178, 255}
	case ColorNameButtonForeground:
		if variant == theme.VariantLight {
			return theme.DefaultTheme().Color(name, variant)
		}
		return color.RGBA{168, 168, 178, 255}
	case ColorNameFocusText:
		if variant == theme.VariantLight {
			return color.RGBA{0, 0, 0, 255}
		}
		return color.RGBA{255, 255, 255, 255}
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m *GophantTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m *GophantTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameInputRadius:
		return 4
	}
	return theme.DefaultTheme().Size(name)
}

func (m *GophantTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}
