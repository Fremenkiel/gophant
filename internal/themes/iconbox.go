package themes

import (
	"fyne.io/fyne/v2"
)

type IconBox struct {
}

func (i *IconBox) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return objects[0].MinSize()
}

func (i *IconBox) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	l := objects[0]
	var ic fyne.CanvasObject
	if len(objects) > 1 {
		ic = objects[1]
	}

	is := fyne.NewSize(0, 0)
	if ic != nil {
		is = ic.MinSize()
		ip := fyne.NewPos(5, (containerSize.Height - is.Height) / 2)
		ic.Resize(is)
		ic.Move(ip)
	}

	ls := l.MinSize()
	lp := fyne.NewPos(10 + is.Width, 0)
	l.Resize(ls)
	l.Move(lp)
}
