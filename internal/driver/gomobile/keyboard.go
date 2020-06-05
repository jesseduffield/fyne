package gomobile

import (
	"github.com/fyne-io/mobile/app"
	"github.com/jesseduffield/fyne"
	"github.com/jesseduffield/fyne/driver/mobile"
)

func showVirtualKeyboard(keyboard mobile.KeyboardType) {
	if driver, ok := fyne.CurrentApp().Driver().(*mobileDriver); ok {
		driver.app.ShowVirtualKeyboard(app.KeyboardType(keyboard))
	}
}

func hideVirtualKeyboard() {
	if driver, ok := fyne.CurrentApp().Driver().(*mobileDriver); ok {
		driver.app.HideVirtualKeyboard()
	}
}
