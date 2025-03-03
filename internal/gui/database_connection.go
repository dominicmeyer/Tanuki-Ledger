package gui

import "fyne.io/fyne/v2/widget"

func NewDatabaseConnectionButton() *widget.Button {
	return widget.NewButton(
		"Add Database Connection",
		func() {},
	)
}
