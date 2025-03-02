package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Tanuki-Ledger")

	w.SetContent(createAddDatabaseButton())

	w.ShowAndRun()
}

func createAddDatabaseButton() *widget.Button {
	button := widget.NewButton("Add Database Connection", func() {})

	return button
}
