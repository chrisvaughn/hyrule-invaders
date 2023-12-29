package engine

import "golang.org/x/image/font"

func RenderedStringLength(str string, fnt font.Face) (int, int) {
	// Measure the size of the rendered string
	bounds, _ := font.BoundString(fnt, str)
	width := (bounds.Max.X - bounds.Min.X).Ceil()
	height := (bounds.Max.Y - bounds.Min.Y).Ceil()

	return width, height
}
