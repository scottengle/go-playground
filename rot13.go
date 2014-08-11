package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Define a new type which has an internal io.Reader
type rot13Reader struct {
	r io.Reader
}

// Implement the Read function, which is equivalent to decoding the string
func (rot13 *rot13Reader) Read(p []byte) (n int, err error) {

	return rot13.Decode(p)
}

// Decode the byte slice, using ROT13
func (rot13 *rot13Reader) Decode(p []byte) (n int, err error) {

	// Read from the internal io.Reader
	num, err := rot13.r.Read(p)

	// Iterate over the byte slice
	for i, elem := range p {

		if (elem >= 'A' && elem < 'N') || (elem >= 'a' && elem < 'n') {

			// A through M (lowercase and uppercase) are rotated forward
			p[i] += 13

		} else if (elem >= 'N' && elem < 'Z') || (elem >= 'n' && elem < 'z') {

			// N through Z (lowercase and uppercase) are rotated backward
			p[i] -= 13

		}
	}
	return num, err
}

func main() {
	encodedString := "Lbh penpxrq gur pbqr!"
	s := strings.NewReader(encodedString)
	r := rot13Reader{s}

	fmt.Println("Encoded String:", encodedString)
	fmt.Print("Decoded String: ")
	io.Copy(os.Stdout, &r)
	fmt.Println()
}
