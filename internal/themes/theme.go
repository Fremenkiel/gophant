package themes

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type GophantTheme struct {}

var _ fyne.Theme = (*GophantTheme)(nil)

func (m *GophantTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	log.Print(name)
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.RGBA{10, 10, 10, 0}
	}

	if name == theme.ColorNameSeparator {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.RGBA{10, 10, 10, 0}
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m *GophantTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m *GophantTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (m *GophantTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}
