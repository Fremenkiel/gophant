package dialogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/interfaces"
)

type ErrorDialog struct {
	Window					fyne.Window
}

var _ interfaces.ErrorReporter = (*ErrorDialog)(nil)

func NewErrorDialog() *ErrorDialog {
	a := fyne.CurrentApp()
	w := a.NewWindow("An error occurred")
	w.Resize(fyne.NewSize(500, 400))

	return &ErrorDialog{Window: w}
}


func (d *ErrorDialog) Report(err error) {
      if err == nil {
                return
        }
        fyne.Do(func() {
                d.Window.SetContent(widget.NewRichTextWithText(err.Error()))
                d.Window.Content().Refresh()
                d.Window.Show()
        })
}
