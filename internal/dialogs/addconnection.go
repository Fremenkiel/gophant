package dialogs

import (
	"errors"
	"image/color"
	"regexp"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/Fremenkiel/gophant/v2/internal/elements"
	"github.com/Fremenkiel/gophant/v2/internal/fs"
	"github.com/Fremenkiel/gophant/v2/internal/interfaces"
	"github.com/Fremenkiel/gophant/v2/internal/models"
	"github.com/Fremenkiel/gophant/v2/internal/th"
)

type AddConnectionDialog struct {
	Window					fyne.Window
	reporter				interfaces.ErrorReporter

	Refresh					func(models.Connection)
}

func NewAddConnectionDialog(r interfaces.ErrorReporter, refresh func(models.Connection)) *AddConnectionDialog {
	a := fyne.CurrentApp()
	t := a.Settings().Theme()
	v := a.Settings().ThemeVariant()

	rAddress := regexp.MustCompile(`[sS]erver=(?P<server>[^;]*);`)
	rDb := regexp.MustCompile(`[dD]atabase=(?P<database>[^;]*);`)
	rPort := regexp.MustCompile(`[pP]ort=(?P<port>[0-9]*);`)
	rUser := regexp.MustCompile(`[uU]ser [iI]d=(?P<userid>[^;]*);`)
	rPass := regexp.MustCompile(`[pP]assword=(?P<password>[^;]*);`)

	eName := elements.NewInput()
	eName.SetValidator(func(i string) error {
		if len(i) == 0 {
			return errors.New("Name is required")
		}
		return nil
	})

	ul := canvas.NewText("URI", color.Transparent)
	ul.Color = t.Color(th.ColorNameLabelText, v)
	ul.TextSize = 11

	ub := elements.NewIconButton(fs.IconNameCopy, nil)

	ui := elements.NewInputAdornment(ul, ub)
	ui.StartSpacer = true
	ui.EndSpacer = true

	eAddress := elements.NewInput()
	eDb := elements.NewInput()
	ePort := elements.NewInput()
	ePort.SetValidator(func(i string) error {
		_, err := strconv.ParseInt(i, 0, 16)
		if err != nil {
			return errors.New("Invalid port")
		}
		return nil
	})
	eUser := elements.NewInputAdornment(canvas.NewText("URI", color.White), nil)
	ePass := elements.NewInput()

	eConnection := elements.NewInput()
	eConnection.SetOnChanged(func(s string) {
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
	})

	//s := widget.NewSeparator()

	w := a.NewWindow("Add Connection")
	/*
	f := &widget.Form{
		Items: []*widget.FormItem{
			{ Widget: eConnection, HintText: "Enter a connection string."},
			{ Widget: s },
			{ Text: "Name", Widget: eName, HintText: "Name is required."},
			{ Text: "Address", Widget: eAddress},
			{ Text: "Database", Widget: eDb},
			{ Text: "Port", Widget: ePort},
			{ Text: "Username", Widget: eUser},
			{ Text: "Password", Widget: ePass},
			{ Text: "Default role", Widget: ePass},
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

			p, err := strconv.ParseInt(ePort.Text(), 0, 16)
			if err != nil {
				r.Report(errors.New("Invalid port"))
				return 
			}

			c := models.Connection{
				ID: &id,
				Name: eName.Text(),
				Address: eAddress.Text(),
				Permission: "ro",
				Port: uint16(p),
				Database: eDb.Text(),
				Username: eUser.Text(),
				Password: ePass.Text(),
				Status: models.OFFLINE,
			}
			
			connections.Set(id.String(), c)
			refresh(c)

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
	*/

	var fit []*elements.FormItem
	fi := elements.NewFormItem(ui, "Host", "resolves via DNS")
	fi.Required = true
	fit = append(fit, fi)
	ft := elements.NewForm(fit)

	w.SetContent(ft)
	w.SetPadded(false)
	w.Resize(fyne.NewSize(500, 400))

	return &AddConnectionDialog{Window: w, Refresh: refresh}
}

func (d *AddConnectionDialog) Open() {
	d.Window.Show()
}

func (d *AddConnectionDialog) Hide() {
	d.Window.Hide()
}
