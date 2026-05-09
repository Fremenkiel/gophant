package th

import "fyne.io/fyne/v2"

type Sidebar struct {
}

func (s *Sidebar) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()

		if childSize.Width > w {
			w = childSize.Width
		}
		h += childSize.Height
	}
	return fyne.NewSize(w, h)
}

func (s *Sidebar) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, 0)
	t := objects[0]
	ts := t.MinSize()
	t.Resize(ts)
	t.Move(pos)

	for _, obj := range objects[1:] {
		pos = fyne.NewPos(0, ts.Height)
		os := fyne.NewSize(containerSize.Width, containerSize.Height - ts.Height)
		obj.Resize(os)
		obj.Move(pos)
	}
}
