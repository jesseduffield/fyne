package playground

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"

	"github.com/jesseduffield/fyne"
	"github.com/jesseduffield/fyne/internal/app"
	"github.com/jesseduffield/fyne/theme"
)

func imageToPlayground(img image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		fyne.LogError("Failed to encode image", err)
		return
	}

	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println("IMAGE:" + enc)
}

// RenderCanvas takes a canvas and converts it into an inline image for showing in the playground
func RenderCanvas(c fyne.Canvas) {
	imageToPlayground(c.Capture())
}

// RenderWindow takes a window and converts it's canvas into an inline image for showing in the playground
func RenderWindow(w fyne.Window) {
	imageToPlayground(w.Canvas().Capture())
}

// Render takes a canvasobject and converts it into an inline image for showing in the playground
func Render(obj fyne.CanvasObject) {
	c := NewSoftwareCanvas()
	c.SetContent(obj)

	fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
	app.ApplyThemeTo(obj, c)
	RenderCanvas(c)
}
