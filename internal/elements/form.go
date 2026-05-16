package elements

import (
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
	r.background.Resize(size)
	r.form.Resize(size)

	p := fyne.NewPos(20, 20)
	for _, obj := range r.form.Items {
		w := size.Width
		s := obj.MinSize()
		if obj.Split {
			w = size.Width / 2
			s = fyne.NewSize(w - 30, s.Height)
		} else {
			s = fyne.NewSize(w - 40, s.Height)
		}
		obj.Resize(s)
		obj.Move(p)
		if obj.Split && p.X == 20 {
			p = fyne.NewPos(size.Width / 2 + 10, p.Y)
			continue
		}
		p = fyne.NewPos(20, p.Y + s.Height + 20)
	}
}

func (r *formRenderer) applyTheme() {
	t := fyne.CurrentApp().Settings().Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	if bg := r.background; bg != nil {
		bg.FillColor = t.Color(theme.ColorNameBackground, v)
		bg.CornerRadius = 0
	}
}

func (r *formRenderer) Refresh() {
	r.applyTheme()
}

