// +build android ios mobile

package app

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jesseduffield/fyne/theme"
)

func TestDefaultTheme(t *testing.T) {
	assert.Equal(t, theme.LightTheme(), defaultTheme())
}
