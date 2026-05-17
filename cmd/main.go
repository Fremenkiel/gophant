package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Fremenkiel/gophant/v2/internal/fragments"
	"github.com/Fremenkiel/gophant/v2/internal/database"
	"github.com/Fremenkiel/gophant/v2/internal/th"
	"github.com/Fremenkiel/gophant/v2/internal/utils"
	"github.com/Fremenkiel/gophant/v2/pkg/dotenv"
)

func main() {
	const envFile = ".env"
	if _, err := os.Stat(envFile); err == nil {
		if err := dotenv.Load(envFile); err != nil {
			log.Printf("load %s:", envFile)
			log.Fatal(err)
		}
	}

	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	database.ApplyMigrations()
	database.SeedData()

	a := app.New()
	a.Settings().SetTheme(&th.GophantTheme{})

	w := a.NewWindow("Main page")
	w.Resize(fyne.NewSize(1000, 800))
	w.SetPadded(false)
	w.SetMainMenu(fyne.NewMainMenu(fyne.NewMenu("test", fyne.NewMenuItem("test.action", func() {}))))
	w.SetMaster()

	ls := fragments.NewMainLayout(w)

	w.SetContent(ls.BuildLayout())
	utils.MapMainKeyBindings(w)

	w.Show()	
	a.Run()
}


