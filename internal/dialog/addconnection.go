package dialog

import (
	"log"
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/elements"
	"github.com/Fremenkiel/gophant/v2/internal/models"
	"github.com/Jipok/go-persist"
	"github.com/google/uuid"
)

type AddConnectionDialog struct {
	Window					fyne.Window
	ConnectionList	*elements.ConnectionList
}
//Server=localhost;Database=bookingboard;Port=5432;User Id=bookingboard;Password=bookingboard;Pooling=false;Trust Server Certificate=true;

func NewAddConnectionDialog(a fyne.App, cl *elements.ConnectionList) *AddConnectionDialog {
	rAddress := regexp.MustCompile(`[sS]erver=(?P<server>[^;]*);`)
	rDb := regexp.MustCompile(`[dD]atabase=(?P<database>[^;]*);`)
	rPort := regexp.MustCompile(`[pP]ort=(?P<port>[0-9]*);`)
	rUser := regexp.MustCompile(`[uU]ser [iI]d=(?P<userid>[^;]*);`)
	rPass := regexp.MustCompile(`[pP]assword=(?P<password>[^;]*);`)

	eName := widget.NewEntry()
	eAddress := widget.NewEntry()
	eDb := widget.NewEntry()
	ePort := widget.NewEntry()
	eUser := widget.NewEntry()
	ePass := widget.NewEntry()

	eConnection := widget.NewEntry()
	eConnection.OnChanged = func(s string) {
		if addressMatch := rAddress.FindStringSubmatch(s); len(addressMatch) == 2 {
			eAddress.SetText(addressMatch[1])
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
			{ Text: "Address", Widget: eAddress},
			{ Text: "Database", Widget: eDb},
			{ Text: "Port", Widget: ePort},
			{ Text: "Username", Widget: eUser},
			{ Text: "Password", Widget: ePass},
		},
		OnSubmit: func() {
			connections, err := persist.OpenSingleMap[models.Connection]("connections.db")
			if err != nil {
				log.Fatal(err)
			}
			defer connections.Store.Close()

			id, err := uuid.NewV7()
			if err != nil {
				log.Fatal(err)
			}
			connections.Set(id.String(), models.Connection{
				ID: &id,
				Name: eName.Text,
				Address: eAddress.Text,
				Port: ePort.Text,
				Database: eDb.Text,
				Username: eUser.Text,
				Password: ePass.Text,
				Status: models.OFFLINE,
			})
			cl.Refresh()

			eConnection.SetText("")
			eName.SetText("")
			eAddress.SetText("")
			ePort.SetText("")
			eDb.SetText("")
			eUser.SetText("")
			ePass.SetText("")
			w.Hide()
		},
	}

	w.SetContent(f)
	w.Resize(fyne.NewSize(500, 400))

	return &AddConnectionDialog{Window: w, ConnectionList: cl}
}


func (d *AddConnectionDialog) Open() {
	d.Window.Show()
}
