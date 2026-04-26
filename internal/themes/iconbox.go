package themes

import (
	"fyne.io/fyne/v2"
)

type IconBox struct {
}

func (i *IconBox) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return objects[1].MinSize()
}

func (i *IconBox) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	ic := objects[0]
	l := objects[1]

	is := ic.MinSize()
	ip := fyne.NewPos(5, (containerSize.Height - is.Height) / 2)
	ic.Resize(is)
	ic.Move(ip)

	ls := l.MinSize()
	lp := fyne.NewPos(10 + is.Width, 0)
	l.Resize(ls)
	l.Move(lp)
}
