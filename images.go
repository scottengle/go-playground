package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

// Define an image struct with width and height properties
// Also allow a custom transformer to be specified.
type Image struct {
	Width       int
	Height      int
	Transformer func(x, y int) uint8
}

// Implement the ColorModel method (required for Image)
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Implement the Bounds method (required for Image)
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

// Implement the At method. This method uses a custom
// transformer to produce the color at the specified
// x and y coordinates.
func (img Image) At(x, y int) color.Color {
	v := img.Transformer(x, y)
	return color.RGBA{v, v, 255, 255}
}

// Implement a custom transformer for the image.
// This function is purposely lowercased.
func transformer(x, y int) uint8 {
	return uint8(x ^ 2 + 2*x - y)
}

func main() {
	// Copy the contents of this file and execute it using
	// the Go Playground at http://play.golang.go

	m := Image{200, 200, transformer}
	pic.ShowImage(m)
}
