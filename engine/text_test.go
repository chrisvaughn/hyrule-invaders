package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/image/font/basicfont"
)

// TestRenderedStringLength tests if we are correctly measuring the rendered string width and height.
// For the basicfont.Face7x13, a string of length 5 should have a width of 35 (since each character is 7 pixels wide)
// and a height of 13 (since that's the height of the font).
func TestRenderedStringLength(t *testing.T) {
	fnt := basicfont.Face7x13
	str := "hello"
	width, height := RenderedStringLength(str, fnt)
	assert.Equal(t, 34, width)
	assert.Equal(t, 13, height)
}
