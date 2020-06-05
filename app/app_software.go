// +build ci nacl

package app

import (
	"github.com/jesseduffield/fyne"
	"github.com/jesseduffield/fyne/internal/painter/software"
	"github.com/jesseduffield/fyne/test"
)

// NewWithID returns a new app instance using the test (headless) driver.
// The ID string should be globally unique to this app.
func NewWithID(id string) fyne.App {
	return newAppWithDriver(test.NewDriverWithPainter(software.NewPainter()), id)
}
