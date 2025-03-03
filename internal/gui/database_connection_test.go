package gui_test

import (
	"testing"

	"github.com/dominicmeyer/Tanuki-Ledger/internal/gui"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnectionButton_NotNil(t *testing.T) {
	t.Parallel()

	button := gui.NewDatabaseConnectionButton()
	assert.NotNil(t, button, "Button shouldn't be nil")
}

func TestDatabaseConnectionButton_Text(t *testing.T) {
	t.Parallel()

	button := gui.NewDatabaseConnectionButton()
	assert.Equal(t, "Add Database Connection", button.Text, "Button Text wasn't right")
}
