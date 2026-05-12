package elements

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type FormItem struct {
	widget.BaseWidget

	input 			fyne.CanvasObject
	label				*canvas.Text
	helpertext	*canvas.Text

}

func NewFormItem(input fyne.CanvasObject, label, helpertext string) *FormItem {
	var l *canvas.Text
	if len(label) > 0 {
		l = canvas.NewText(label, color.Transparent)
	}

	var h *canvas.Text
	if len(helpertext) > 0 {
		h = canvas.NewText(helpertext, color.Transparent)
	}

	c := &FormItem{
		input: input,
		label: l,
		helpertext: h,
	}
	c.ExtendBaseWidget(c)
	return c
}

func (f *FormItem) CreateRenderer() fyne.WidgetRenderer {
	f.ExtendBaseWidget(f)
	t := f.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	background := canvas.NewRectangle(t.Color(theme.ColorNameButton, v))
	background.CornerRadius = t.Size(theme.SizeNameInputRadius)

	objects := []fyne.CanvasObject{
		background,
		f.input,
	}

	if f.label != nil {
		objects = append(objects, f.label)
	}

	if f.helpertext != nil {
		objects = append(objects, f.helpertext)
	}

	r := &formItemRenderer{
		BaseRenderer: NewBaseRenderer(objects),
		item: f,
		background: background,
	}
	r.applyTheme()
	return r
}

type formItemRenderer struct {
	BaseRenderer

	background	*canvas.Rectangle
	item				*FormItem
	layout			fyne.Layout
}

func (r *formItemRenderer) MinSize() fyne.Size {
	return r.item.input.MinSize()
}

func (r *formItemRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	r.item.input.Resize(size)
}

func (r *formItemRenderer) applyTheme() {
	t := fyne.CurrentApp().Settings().Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	if bg := r.background; bg != nil {
		bg.StrokeColor = t.Color(theme.ColorNameSeparator, v)
		bg.StrokeWidth = 1
		bg.CornerRadius = 0
	}
}

func (r *formItemRenderer) Refresh() {
	r.applyTheme()
}

