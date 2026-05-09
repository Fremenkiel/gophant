package layouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

func NewBorderBox() fyne.Layout {
	return &borderBoxLayout{
		paddingFunc: theme.Padding,
	}
}

/*
func (c *BorderBox) CreateRenderer() fyne.WidgetRenderer {
	c.ExtendBaseWidget(c)
	t := c.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	background := canvas.NewRectangle(t.Color(theme.ColorNameButton, v))
	background.CornerRadius = t.Size(theme.SizeNameInputRadius)

	objects := []fyne.CanvasObject{
		background,
	}
	for _, obj := range c.elements {
		objects = append(objects, obj)
	}

	r := &borderBoxLayout{
		BaseRenderer: elements.NewBaseRenderer(objects),
		elements: c.elements,
		background: background,
	}
	r.applyTheme()
	return r
}
*/
var _ fyne.Layout = (*borderBoxLayout)(nil)

type borderBoxLayout struct {
	paddingFunc func() float32
}

func (b *borderBoxLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	visibleObjects := 0
	// Size taken up by visible objects
	total := float32(0)

	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		visibleObjects++
		total += child.MinSize().Height
	}

	padding := b.paddingFunc()

	// Spacers split extra space equally
	x, y := float32(0), float32(0)
	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		child.Move(fyne.NewPos(x, y))

		height := child.MinSize().Height
		y += padding + height
		child.Resize(fyne.NewSize(size.Width, height))
	}
}

// MinSize finds the smallest size that satisfies all the child objects.
// For a BoxLayout this is the width of the widest item and the height is
// the sum of all children combined with padding between each.
func (b *borderBoxLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	minSize := fyne.NewSize(0, 0)
	addPadding := false
	padding := b.paddingFunc()
	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		childMin := child.MinSize()
		minSize.Width = fyne.Max(childMin.Width, minSize.Width)
		minSize.Height += childMin.Height
		if addPadding {
			minSize.Height += padding
		}
		addPadding = true
	}
	return minSize
}

/*
func (r *borderBoxLayout) MinSize() fyne.Size {
		return fyne.NewSize(0, 0)
}

func (r *borderBoxLayout) Layout(size fyne.Size) {
	r.background.Resize(size)

	y := float32(0)
	for _, obj := range r.elements {
		s := obj.MinSize()
		obj.Resize(s)
		obj.Move(fyne.NewPos(0, y))
		y = y + s.Height 
	}
}

*/
