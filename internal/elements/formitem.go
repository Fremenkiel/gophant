package elements

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/th"
)

type FormItem struct {
	widget.BaseWidget

	input 			fyne.CanvasObject
	label				*canvas.Text
	helpertext	*canvas.Text

	Required, Split		bool

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
	
	ri := canvas.NewText("*", color.Transparent)
	ri.Hide()

	objects := []fyne.CanvasObject{
		background,
		ri,
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
		requiredIndicator: ri,
		background: background,
	}
	r.applyTheme()
	return r
}

type formItemRenderer struct {
	BaseRenderer

	background	*canvas.Rectangle
	item				*FormItem
	requiredIndicator	*canvas.Text
	layout			fyne.Layout
}

func (r *formItemRenderer) MinSize() fyne.Size {
	is := r.item.input.MinSize()
	lh := r.item.label.MinSize().Height
	return fyne.NewSize(is.Width, is.Height + lh + 5)
}

func (r *formItemRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	p := fyne.NewPos(0, 0)

	lh := float32(0)
	if l := r.item.label; l != nil {
		ls := l.MinSize()
		l.Resize(ls)
		l.Move(p)
		lh = ls.Height

		r.requiredIndicator.Resize(r.requiredIndicator.MinSize())
		r.requiredIndicator.Move(fyne.NewPos(ls.Width + 5, 0))
	}

	if h := r.item.helpertext; h != nil {
		hs := h.MinSize()
		h.Resize(hs)
		h.Move(fyne.NewPos(size.Width - hs.Width, 0))
	}

	p = p.AddXY(0, lh + 5)
	is := r.item.input.MinSize()
	r.item.input.Move(p)
	r.item.input.Resize(fyne.NewSize(size.Width, is.Height))
}

func (r *formItemRenderer) applyTheme() {
	t := fyne.CurrentApp().Settings().Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	if bg := r.background; bg != nil {
		bg.FillColor = color.Transparent
	}

	if label := r.item.label; label != nil {
		label.Color = t.Color(th.ColorNameLabelText, v)
		label.TextSize = 11
	}

	if helpertext := r.item.helpertext; helpertext != nil {
		helpertext.Color = t.Color(th.ColorNameLabelText, v)
		helpertext.TextSize = 11
	}

	if ri := r.requiredIndicator; ri != nil {
		ri.Color = t.Color(theme.ColorNameError, v)
		ri.TextSize = 11
		if r.item.Required {
			ri.Show()
		}
	}
}

func (r *formItemRenderer) Refresh() {
	r.applyTheme()
}

