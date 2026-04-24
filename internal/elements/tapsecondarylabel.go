package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TapSecondaryLabel struct {
   *widget.Label

    OnTappedSecondary func(pe *fyne.PointEvent)
}

func NewTapSecondaryLabel(text string, tappedRight func(pe *fyne.PointEvent)) *TapSecondaryLabel {
   return &TapSecondaryLabel{
      widget.NewLabel(text),
      tappedRight,
   }
}

func (mc *TapSecondaryLabel) TappedSecondary(pe *fyne.PointEvent) {
    if mc.OnTappedSecondary != nil {
        mc.OnTappedSecondary(pe)
    }
}
