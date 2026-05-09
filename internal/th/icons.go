package th

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/Fremenkiel/gophant/v2/internal/svg"
)

var _ fyne.ThemedResource = (*ThemedResource)(nil)

type ThemedResource struct {
	source fyne.Resource

	FillName fyne.ThemeColorName
	StrokeName fyne.ThemeColorName
}

func NewColoredResource(src fyne.Resource, fillName, strokeName fyne.ThemeColorName) *ThemedResource {
	return &ThemedResource{
		source:    src,
		FillName: fillName,
		StrokeName: strokeName,
	}
}

func (res *ThemedResource) ThemeColorName() fyne.ThemeColorName {
	if res.FillName != "" {
		return res.FillName
	}

	return theme.ColorNameBackground
}

func (res *ThemedResource) ThemeColorStrokeName() fyne.ThemeColorName {
	if res.StrokeName != "" {
		return res.StrokeName
	}

	return theme.ColorNameForeground
}

func (res *ThemedResource) Name() string {
	return string(res.ThemeColorName()) + "_" + string(res.ThemeColorStrokeName()) + "_" + unwrapResource(res.source).Name()
}

func (res *ThemedResource) Content() []byte {
	return colorizeLogError(unwrapResource(res.source).Content(), theme.Color(res.ThemeColorName()), theme.Color(res.ThemeColorStrokeName()))
}

func unwrapResource(res fyne.Resource) fyne.Resource {
	for {
		switch typedRes := res.(type) {
		case *ThemedResource:
			res = typedRes.source
		default:
			return res
		}
	}
}

func colorizeLogError(src []byte, fClr, sClr color.Color) []byte {
	content, err := svg.ColorizeStroke(src, sClr)
	if err != nil {
		fyne.LogError("", err)
	}

	content, err = svg.Colorize(content, fClr)
	if err != nil {
		fyne.LogError("", err)
	}
	return content
}
