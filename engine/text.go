package engine

import "golang.org/x/image/font"

// RenderedStringLength measures the size of the rendered string using the given font face.
// It returns the width and height as integers.
func RenderedStringLength(str string, fnt font.Face) (int, int) {
	// Measure the size of the rendered string
	bounds, _ := font.BoundString(fnt, str)
	width := (bounds.Max.X - bounds.Min.X).Ceil()
	height := (bounds.Max.Y - bounds.Min.Y).Ceil()

	return width, height
}
