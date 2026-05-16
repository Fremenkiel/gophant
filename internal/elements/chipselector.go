package elements

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/fs"
)

type ChipSelector struct {
	widget.BaseWidget
	chips		[]*Chip
	
	selected	int

	AddChip	func(*fyne.PointEvent)
}

func NewChipSelector(addChip func(*fyne.PointEvent), chips ...*Chip) *ChipSelector {
	b := &ChipSelector{chips: chips, selected: -1, AddChip: addChip}

	for index, obj := range b.chips {
		obj.OnTapped = func(pe *fyne.PointEvent) {
			b.SetSelected(index)
		}
	}

	b.ExtendBaseWidget(b)
	return b
}

func (i *ChipSelector) CreateRenderer() fyne.WidgetRenderer {
	i.ExtendBaseWidget(i)
	th := i.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	background := canvas.NewRectangle(th.Color(theme.ColorNameButton, v))
	background.CornerRadius = th.Size(theme.SizeNameInputRadius)

	objects := []fyne.CanvasObject{
		background,
	}

	for _, obj := range i.chips {
		objects = append(objects, obj)
	}

	var addChip *IconChip
	if i.AddChip != nil {
		addChip = NewIconChip("new", fs.IconNameAdd, i.AddChip, nil)
		objects = append(objects, addChip)
	}

	b := &chipSelectorRenderer{
		BaseRenderer: NewBaseRenderer(objects),
		selector: i,
		addChip: addChip,
		background: background,
	}
	b.applyTheme()
	return b
}

func (i *ChipSelector) SetSelected(index int) {
	if i.selected != -1 {
		i.chips[i.selected].SetFocus(false)
		i.chips[i.selected].Refresh()
	}

	if i.selected != index {
		i.chips[index].SetFocus(true)
		i.chips[index].Refresh()
		i.selected = index
	} else {
		i.selected = -1
	}
}

/*
func (i *ChipSelector) SetContent(text string, icon fyne.CanvasObject) {
	i.label.SetText(text)

	if icon != nil {
		icon.Resize(fyne.NewSize(12, 12))
			i.icon.Objects[0] = icon
	}

	i.Refresh()
}
*/

type chipSelectorRenderer struct {
	BaseRenderer

	background *canvas.Rectangle
	selector     *ChipSelector
	addChip			*IconChip
	layout     fyne.Layout
}

func (r *chipSelectorRenderer) MinSize() fyne.Size {
	if chips := r.selector.chips; len(chips) > 0 {
		s := fyne.NewSize(0, chips[0].MinSize().Height)
		for _, obj := range chips {
			s = s.AddWidthHeight(obj.MinSize().Width, s.Height)
		}
		s = s.AddWidthHeight(float32(5 * len(chips)), 0)
		if r.addChip != nil {
			s = s.AddWidthHeight(r.addChip.MinSize().Width + 5, 0)
		}
		return s
	}
	return fyne.NewSize(0, 0)
}

func (r *chipSelectorRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)

	p := fyne.NewPos(0, 0)
	for _, obj := range r.selector.chips {
		s := obj.MinSize()
		obj.Resize(s)
		obj.Move(p)

		p = p.AddXY(s.Width + 5, 0)
	}
	if addChip := r.addChip; addChip != nil {
		s := addChip.MinSize()
		addChip.Resize(s)
		addChip.Move(p)
	}
}

func (r *chipSelectorRenderer) applyTheme() {
	/*
	v := fyne.CurrentApp().Settings().ThemeVariant()
	t := fyne.CurrentApp().Settings().Theme() 
*/
	if bg := r.background; bg != nil {
		bg.FillColor = color.Color(color.Transparent)
		bg.Refresh()
	}
}

func (r *chipSelectorRenderer) Refresh() {
	r.applyTheme()
}

/*
func (r *chipSelectorRenderer) buttonColorNames() (foreground, background, backgroundBlend fyne.ThemeColorName) {
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
*/
