// +build !windows,!android,!ios

package dialog

import (
	"github.com/jesseduffield/fyne"
	"github.com/jesseduffield/fyne/widget"
)

func (f *fileDialog) loadPlaces() []fyne.CanvasObject {
	return []fyne.CanvasObject{widget.NewButton("Computer", func() {
		f.setDirectory("/")
	})}
}

func isHidden(file, _ string) bool {
	return len(file) == 0 || file[0] == '.'
}

func fileOpenOSOverride(*FileDialog) bool {
	return false
}

func fileSaveOSOverride(*FileDialog) bool {
	return false
}
