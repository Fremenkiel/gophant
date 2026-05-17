package dialogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func ReportError(err error) {
	if err == nil {
		return
	}

	fyne.Do(func() {
		a := fyne.CurrentApp()
		w := a.NewWindow("An error occurred")
		w.Resize(fyne.NewSize(500, 400))

		w.SetContent(widget.NewRichTextWithText(err.Error()))
		w.Content().Refresh()
		w.Show()
	})
}
