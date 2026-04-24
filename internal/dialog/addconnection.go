package dialog

import (
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type AddConnectionDialog struct {
	Window				fyne.Window
}
//Server=localhost;Database=bookingboard;Port=5432;User Id=bookingboard;Password=bookingboard;Pooling=false;Trust Server Certificate=true;

func NewAddConnectionDialog(a fyne.App) *AddConnectionDialog {
	rIp := regexp.MustCompile(`[sS]erver=(?P<server>[^;]*);`)
	rDb := regexp.MustCompile(`[dD]atabase=(?P<database>[^;]*);`)
	rPort := regexp.MustCompile(`[pP]ort=(?P<port>[0-9]*);`)
	rUser := regexp.MustCompile(`[uU]ser [iI]d=(?P<userid>[^;]*);`)
	rPass := regexp.MustCompile(`[pP]assword=(?P<password>[^;]*);`)

	eName := widget.NewEntry()
	eIp := widget.NewEntry()
	eDb := widget.NewEntry()
	ePort := widget.NewEntry()
	eUser := widget.NewEntry()
	ePass := widget.NewEntry()

	eConnection := widget.NewEntry()
	eConnection.OnChanged = func(s string) {
		if ipMatch := rIp.FindStringSubmatch(s); len(ipMatch) == 2 {
			eIp.SetText(ipMatch[1])
		}
		if dbMatch := rDb.FindStringSubmatch(s); len(dbMatch) == 2 {
			eDb.SetText(dbMatch[1])
		}
		if portMatch := rPort.FindStringSubmatch(s); len(portMatch) == 2 {
			ePort.SetText(portMatch[1])
		}
		if userMatch := rUser.FindStringSubmatch(s); len(userMatch) == 2 {
			eUser.SetText(userMatch[1])
		}
		if passMatch := rPass.FindStringSubmatch(s); len(passMatch) == 2 {
			ePass.SetText(passMatch[1])
		}
	}

	s := widget.NewSeparator()

	w := a.NewWindow("Add Connection")
	f := &widget.Form{
		Items: []*widget.FormItem{
			{ Text: "Connection", Widget: eConnection},
			{ Widget: s },
			{ Text: "Name", Widget: eName},
			{ Text: "Ip", Widget: eIp},
			{ Text: "Database", Widget: eDb},
			{ Text: "Port", Widget: ePort},
			{ Text: "Username", Widget: eUser},
			{ Text: "Password", Widget: ePass},
		},
	}

	w.SetContent(f)
	w.Resize(fyne.NewSize(500, 400))

	return &AddConnectionDialog{Window: w}
}


func (d *AddConnectionDialog) Open() {
	d.Window.Show()
}
