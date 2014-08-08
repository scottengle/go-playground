#Golang Playground

This repo is intended to be a collection of Go examples that can be independently run using:

	go run <filename>

##Note

You will see errors regarding main being redeclared when using go get on this repo. You can safely ignore this warning.

##Examples

Filename | Description
--- | ---
fibonacci.go | Calculates the value of a fibonacci sequence using a closure.
newton_raphson_sqrt.go | Approximates the value of a square root using the Newton-Raphson method.
visual_slices.go | Generates a bluescale image using the code.google.com/p/go-tour/pic package. To see the image, run the contents of this file at the Go Playground: http://play.golang.org
word_count.go | Determines word count of a string utilizing maps.