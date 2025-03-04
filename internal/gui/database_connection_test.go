package gui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnectionButton_Text(t *testing.T) {
	t.Parallel()

	button := NewDatabaseConnectionButton()

	assert.Equal(t, "Add Database Connection", button.Text, "Button Text wasn't right")
}

func TestDatabaseConnectionDialog_ItemsCount(t *testing.T) {
	t.Parallel()

	form := NewDatabaseConnectionForm()

	assert.Len(t, form.Items, 4, "Form should have four elements")
}

func TestDatabaseConnectionDialog_ItemsContainStructInputs(t *testing.T) {
	t.Parallel()

	form := NewDatabaseConnectionForm()

	assert.Equal(
		t,
		form.databaseURLInput,
		form.Items[0].Widget,
		"databaseURLInput is not the same as in the base form",
	)
	assert.Equal(
		t,
		form.databaseUsernameInput,
		form.Items[1].Widget,
		"databaseUsernameInput is not the same as in the base form",
	)
	assert.Equal(
		t,
		form.databasePasswordInput,
		form.Items[2].Widget,
		"databasePasswordInput is not the same as in the base form",
	)
	assert.Equal(
		t,
		form.databaseNameInput,
		form.Items[3].Widget,
		"databaseNameInput is not the same as in the base form",
	)
}

func TestDatabaseConnectionDialog_ItemsText(t *testing.T) {
	t.Parallel()

	form := NewDatabaseConnectionForm()

	assert.Equal(t, "url", form.Items[0].Text, "url text wasn't right")
	assert.Equal(t, "username", form.Items[1].Text, "username text wasn't right")
	assert.Equal(t, "password", form.Items[2].Text, "password text wasn't right")
	assert.Equal(t, "database name", form.Items[3].Text, "database name text wasn't right")
}
