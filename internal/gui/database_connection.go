package gui

import (
	"fyne.io/fyne/v2/widget"
)

func NewDatabaseConnectionButton() *widget.Button {
	return widget.NewButton(
		"Add Database Connection",
		nil,
	)
}

type DatabaseConnectionForm struct {
	*widget.Form
	databaseURLInput      *widget.Entry
	databaseUsernameInput *widget.Entry
	databasePasswordInput *widget.Entry
	databaseNameInput     *widget.Entry
}

func NewDatabaseConnectionForm() *DatabaseConnectionForm {
	databaseURLInput := widget.NewEntry()
	databaseUsernameInput := widget.NewEntry()
	databasePasswordInput := widget.NewEntry()
	databaseNameInput := widget.NewEntry()

	form := widget.NewForm(
		widget.NewFormItem("url", databaseURLInput),
		widget.NewFormItem("username", databaseUsernameInput),
		widget.NewFormItem("password", databasePasswordInput),
		widget.NewFormItem("database name", databaseNameInput),
	)
	form.CancelText = "Cancel"
	form.SubmitText = "Connect"

	databaseConnectionForm := DatabaseConnectionForm{
		form,
		databaseURLInput,
		databaseUsernameInput,
		databasePasswordInput,
		databaseNameInput,
	}

	return &databaseConnectionForm
}
