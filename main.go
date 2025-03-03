package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/dominicmeyer/Tanuki-Ledger/internal/gui"
)

func main() {
	a := app.New()
	w := a.NewWindow("Tanuki-Ledger")

	w.SetContent(gui.NewDatabaseConnectionButton())

	w.ShowAndRun()
}
