package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Creates a slice of dy slices each of length dx
	// containing an 8-bit unsigned integer.

	// Initialize the outer slice
	y := make([][]uint8, dy)

	for i := range y {

		// Iterate over the outer slice to initialize the inner slice
		y[i] = make([]uint8, dx)

		// Assign calculated uint8 values for each inner slice position
		for j := range y[i] {

			// using other formulas here will change the resulting image
			y[i][j] = uint8(i ^ 2 + 2*i - j)

		}
	}
	return y
}

func main() {

	// Copy the contents of this file and execute it using
	// the Go Playground at http://play.golang.go

	pic.Show(Pic)
}
