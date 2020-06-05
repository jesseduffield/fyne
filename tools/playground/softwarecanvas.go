package playground

import (
	"github.com/jesseduffield/fyne/internal/painter/software"
	"github.com/jesseduffield/fyne/test"
)

// NewSoftwareCanvas creates a new canvas in memory that can render without hardware support
func NewSoftwareCanvas() test.WindowlessCanvas {
	return test.NewCanvasWithPainter(software.NewPainter())
}
