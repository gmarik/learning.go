// Package image defines the Image interface:
//
// package image
//
// type Image interface {
//     ColorModel() color.Model
//     Bounds() Rectangle
//     At(x, y int) color.Color
// }
// (See the documentation for all the details.)
//
// Also, color.Color and color.Model are interfaces, but we'll ignore that by using the predefined implementations color.RGBA and color.RGBAModel.
package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
