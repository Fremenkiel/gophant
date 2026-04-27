package dialogs

import (
	"errors"
	"regexp"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Fremenkiel/gophant/v2/internal/fragments"
	"github.com/Fremenkiel/gophant/v2/internal/interfaces"
	"github.com/Fremenkiel/gophant/v2/internal/models"
	"github.com/Jipok/go-persist"
	"github.com/google/uuid"
)

type AddConnectionDialog struct {
	Window					fyne.Window
	ConnectionList	*fragments.ConnectionList
	reporter				interfaces.ErrorReporter
}

func NewAddConnectionDialog(a fyne.App, r interfaces.ErrorReporter, cl *fragments.ConnectionList) *AddConnectionDialog {
	rAddress := regexp.MustCompile(`[sS]erver=(?P<server>[^;]*);`)
	rDb := regexp.MustCompile(`[dD]atabase=(?P<database>[^;]*);`)
	rPort := regexp.MustCompile(`[pP]ort=(?P<port>[0-9]*);`)
	rUser := regexp.MustCompile(`[uU]ser [iI]d=(?P<userid>[^;]*);`)
	rPass := regexp.MustCompile(`[pP]assword=(?P<password>[^;]*);`)

	eName := widget.NewEntry()
	eName.Validator = func(i string) error {
		if len(i) == 0 {
			return errors.New("Name is required")
		}
		return nil
	}
	eAddress := widget.NewEntry()
	eDb := widget.NewEntry()
	ePort := widget.NewEntry()
	ePort.Validator = func(i string) error {
		_, err := strconv.ParseInt(i, 0, 16)
		if err != nil {
			return errors.New("Invalid port")
		}
		return nil
	}
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
			{ Text: "Connection", Widget: eConnection, HintText: "Enter a connection string."},
			{ Widget: s },
			{ Text: "Name", Widget: eName, HintText: "Name is required."},
			{ Text: "Address", Widget: eAddress},
			{ Text: "Database", Widget: eDb},
			{ Text: "Port", Widget: ePort},
			{ Text: "Username", Widget: eUser},
			{ Text: "Password", Widget: ePass},
		},
		OnSubmit: func() {
			connections, err := persist.OpenSingleMap[models.Connection]("connections.db")
			if err != nil {
				r.Report(err)
				return 
			}
			defer connections.Store.Close()

			id, err := uuid.NewV7()
			if err != nil {
				r.Report(err)
				return 
			}

			p, err := strconv.ParseInt(ePort.Text, 0, 16)
			if err != nil {
				r.Report(errors.New("Invalid port"))
				return 
			}
			
			connections.Set(id.String(), models.Connection{
				ID: &id,
				Name: eName.Text,
				Address: eAddress.Text,
				Port: uint16(p),
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
