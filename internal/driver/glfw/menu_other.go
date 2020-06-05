// +build !darwin no_native_menus

package glfw

import "github.com/jesseduffield/fyne"

func hasNativeMenu() bool {
	return false
}

func setupNativeMenu(_ *window, _ *fyne.MainMenu) {
	// no-op
}
