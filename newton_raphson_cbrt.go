package main

import "fmt"
import "math/cmplx"

// Calculate the approximate value of the cube root
// of x using the Newton-Raphson method.
// Returns the complex result and number of iterations
// required to calculate it.
func Cbrt(x complex128) (complex128, int) {

	// Note: This method doesn't work with negative values of x

	// Special case for x = 0
	if x == 0 {
		return 0, 1
	}

	// Keep track of iterations
	i := 0

	// Make a local copy of x
	var z complex128 = x

	// Internal variable to calculate deltas between iterations
	last := z

	for {
		z = z - ((z*z*z - x) / (3 * z * z))

		// Increment the iteration counter
		i++

		// Exit out when the delta of change is sufficiently low
		if cmplx.Abs(z-last) < 10e-16 {
			break
		}

		last = z
	}

	return z, i
}

func main() {

	fmt.Println("Approximate the complex cube root of a number using the Newton–Raphson method.\n")

	var num complex128 = 2
	result, i := Cbrt(num)

	fmt.Printf("Cbrt%v, Approximated in %v iterations\n", num, i)
	fmt.Printf("Approx: %v\n", result)
	fmt.Printf("Actual: %v\n", cmplx.Pow(num, 1.0/3.0))
}
