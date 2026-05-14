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
	ColorNameButtonHover = "ColorNameButtonHover"
	ColorNameFocusText = "ColorNameFocusText"
	ColorNameTransparent = "ColorNameTransparent"
	ColorNameLabelText = "ColorNameLabelText"
)

func (m *GophantTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNamePrimary:
		if variant == theme.VariantLight {
			return theme.DefaultTheme().Color(name, variant)
		}
		return color.RGBA{42, 69, 89, 255}
	case theme.ColorNameInputBorder:
		if variant == theme.VariantLight {
			return theme.DefaultTheme().Color(name, variant)
		}
		return color.RGBA{34, 34, 43, 255}
	case theme.ColorNameBackground:
		if variant == theme.VariantLight {
			return theme.DefaultTheme().Color(name, variant)
		}
		return color.RGBA{19, 19, 24, 255}
	case theme.ColorNameForeground:
		if variant == theme.VariantLight {
			return theme.DefaultTheme().Color(name, variant)
		}
		return color.RGBA{168, 168, 178, 255}
	case theme.ColorNameFocus:
		if variant == theme.VariantLight {
			return theme.DefaultTheme().Color(name, variant)
		}
		return color.RGBA{24, 24, 31, 255}
	case theme.ColorNameHover:
		if variant == theme.VariantLight {
			return theme.DefaultTheme().Color(name, variant)
		}
		return color.RGBA{20, 20, 26, 255}
	case theme.ColorNameSeparator:
		if variant == theme.VariantLight {
			return theme.DefaultTheme().Color(name, variant)
		}
		return color.RGBA{26, 26, 33, 255}
	case theme.ColorNameButton:
		if variant == theme.VariantLight {
			return theme.DefaultTheme().Color(name, variant)
		}
		return color.RGBA{13, 13, 16, 255}
	case theme.ColorNameError:
		if variant == theme.VariantLight {
			return theme.DefaultTheme().Color(name, variant)
		}
		return color.RGBA{226, 107, 99, 255}
	case theme.ColorNameInputBackground:
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
	case ColorNameButtonHover:
		if variant == theme.VariantLight {
			return theme.DefaultTheme().Color(name, variant)
		}
		return color.RGBA{19, 19, 24, 255}
	case ColorNameFocusText:
		if variant == theme.VariantLight {
			return color.RGBA{0, 0, 0, 255}
		}
		return color.RGBA{255, 255, 255, 255}
	case ColorNameLabelText:
		if variant == theme.VariantLight {
			return color.RGBA{0, 0, 0, 255}
		}
		return color.RGBA{108, 108, 120, 255}
	case ColorNameTransparent:
		return color.Transparent
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
