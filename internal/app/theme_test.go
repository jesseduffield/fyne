package app_test

import (
	"testing"

	"github.com/jesseduffield/fyne/internal/app"
	"github.com/jesseduffield/fyne/test"
)

func TestApplySettings_BeforeContentSet(t *testing.T) {
	a := test.NewApp()
	_ = a.NewWindow("NoContent")

	app.ApplySettings(a.Settings(), a)
}
