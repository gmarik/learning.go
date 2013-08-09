// Remember the picture generator you wrote earlier? Let's write another one, but this time it will return an implementation of image.Image instead of a slice of data.
//
// Define your own Image type, implement the necessary methods, and call pic.ShowImage.
//
// Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
//
// ColorModel should return color.RGBAModel.
//
// At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.

package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Width, Height int
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (i *Image) At(x, y int) color.Color {
	return color.RGBA{128 + uint8(x), 128 + uint8(y), 255, 255}
}

func main() {
	m := &Image{100, 100}
	pic.ShowImage(m)
}
