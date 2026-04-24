package themes

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
	toolbar := objects[0]
	size := toolbar.MinSize()
	toolbar.Resize(size)
	toolbar.Move(pos)

	pos = pos.Add(fyne.NewPos(0, size.Height))
	size = fyne.NewSize(containerSize.Width, containerSize.Height - size.Height)
	scroll := objects[1]
	scroll.Resize(size)
	scroll.Move(pos)
}
