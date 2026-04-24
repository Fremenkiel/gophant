package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TapLabel struct {
   *widget.Label

    OnTapped func()
    OnTappedSecondary func(pe *fyne.PointEvent)
}

func NewTapLabel(text string, tappedLeft func(), tappedRight func(pe *fyne.PointEvent)) *TapLabel {
   return &TapLabel{
      widget.NewLabel(text),
      tappedLeft, tappedRight,
   }
}

func (mc *TapLabel) TappedSecondary(pe *fyne.PointEvent) {
    if mc.OnTappedSecondary != nil {
        mc.OnTappedSecondary(pe)
    }
}

func (mc *TapLabel) Tapped(pe *fyne.PointEvent) {
    if mc.OnTapped != nil {
        mc.OnTapped()
    }
}

