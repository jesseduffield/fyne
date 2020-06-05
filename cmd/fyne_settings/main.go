package main

import (
	"github.com/jesseduffield/fyne"
	"github.com/jesseduffield/fyne/app"
	"github.com/jesseduffield/fyne/cmd/fyne_settings/settings"
	"github.com/jesseduffield/fyne/widget"
)

func main() {
	s := settings.NewSettings()

	a := app.New()
	w := a.NewWindow("Fyne Settings")

	appearance := s.LoadAppearanceScreen(w)
	tabs := widget.NewTabContainer(
		&widget.TabItem{Text: "Appearance", Icon: s.AppearanceIcon(), Content: appearance})
	tabs.SetTabLocation(widget.TabLocationLeading)
	w.SetContent(tabs)

	w.Resize(fyne.NewSize(480, 320))
	w.ShowAndRun()
}
