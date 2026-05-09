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
	t.Resize(fyne.NewSize(containerSize.Width, ts.Height))
	t.Move(pos)

	cPos := fyne.NewPos(0, ts.Height)
	for _, obj := range objects[1:] {
		s := obj.MinSize()
		obj.Resize(fyne.NewSize(containerSize.Width, s.Height))
		obj.Move(cPos)
		cPos = fyne.NewPos(0, pos.Y + s.Height)
	}
}
