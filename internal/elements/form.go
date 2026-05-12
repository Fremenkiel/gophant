package elements

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Form struct {
	widget.BaseWidget

	Items []*FormItem
}

func NewForm(items []*FormItem) *Form {
	c := &Form{
		Items: items,
	}
	c.ExtendBaseWidget(c)
	return c
}

func (f *Form) CreateRenderer() fyne.WidgetRenderer {
	f.ExtendBaseWidget(f)
	t := f.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	background := canvas.NewRectangle(t.Color(theme.ColorNameButton, v))
	background.CornerRadius = t.Size(theme.SizeNameInputRadius)


	objects := []fyne.CanvasObject{
		background,
	}

	for _, obj := range f.Items {
		objects = append(objects, obj)
	}

	r := &formRenderer{
		BaseRenderer: NewBaseRenderer(objects),
		form: f,
		background: background,
	}
	r.applyTheme()
	return r
}

type formRenderer struct {
	BaseRenderer

	background	*canvas.Rectangle
	form				*Form
	layout			fyne.Layout
}

func (r *formRenderer) MinSize() fyne.Size {
	return fyne.NewSize(40, 40)
}

func (r *formRenderer) Layout(size fyne.Size) {
	log.Println(size)
	r.background.Resize(size)
	r.form.Resize(size)
	for _, obj := range r.form.Items {
		s := obj.MinSize()
		obj.Resize(fyne.NewSize(size.Width, s.Height))
		obj.Move(fyne.NewPos(0, 0))
	}
}

func (r *formRenderer) applyTheme() {
	t := fyne.CurrentApp().Settings().Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	if bg := r.background; bg != nil {
		bg.StrokeColor = t.Color(theme.ColorNameSeparator, v)
		bg.StrokeWidth = 1
		bg.CornerRadius = 0
	}
}

func (r *formRenderer) Refresh() {
	r.applyTheme()
}

