# Golang Playground

This repo is intended to be a collection of Go examples that can be independently run using:

	go run <filename>

### Note

You will see errors regarding main being redeclared when using go get on this repo. You can safely ignore this warning.

### Files

Filename | Description
--- | ---
fibonacci.go | Calculates the value of a fibonacci sequence using a closure.
http_handlers.go | A simple HTTP handler example to demonstrate how to handle HTTP requests and register routes.<br/><br/>Start this example with:<br/>go run http_handlers.go<br/><br/>Output can be seen in a web browser by opening:<br/>http://localhost:4000/string<br/>http://localhost:4000/struct
newton_raphson_cbrt.go | Approximates the value of a complex cube root using the Newton-Raphson method.
newton_raphson_sqrt.go | Approximates the value of a square root using the Newton-Raphson method.
newton_raphson_sqrt_2.go | Approximates the value of a square root using the Newton-Raphson method. This version throws a proper error when called with negative numbers.
visual_slices.go | Generates a bluescale image using the code.google.com/p/go-tour/pic package. To see the image, run the contents of this file at the Go Playground: http://play.golang.org
word_count.go | Determines word count of a string utilizing maps.