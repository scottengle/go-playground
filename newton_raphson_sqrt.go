package main

import (
	"fmt"
	"math"
)

// Calculate the approximate value of the square root of x.
// By using unsigned integers, I'm sidestepping the case
// where the result is an imaginary number.
func Sqrt(x uint64) (float64, int) {

	// Special case for x = 0
	if x == 0 {
		return 0, 1
	}

	// Keep track of iterations
	i := 0

	// Start with 1 (no dividing by zero!)
	z := float64(1)

	// Convert incoming unsigned integer to a float for calculations
	y := float64(x)

	// Internal variable to calculate deltas between iterations
	last := z

	for {
		z = z - (((z * z) - y) / (2 * z))

		// Increment the iteration counter
		i++

		// Exit out when the delta of change is sufficiently low
		if math.Abs(z-last) < 10e-16 {
			break
		}

		last = z
	}

	return z, i
}

func main() {
	fmt.Println("Approximate the square root of a number using the Newton–Raphson method.\n")

	var num uint64 = 0

	for i := 0; i <= 50; i++ {
		result, iter := Sqrt(num)
		fmt.Printf("Sqrt(%v), Approximated in %v Iterations\n", num, iter)
		fmt.Printf("Approx: %v\n", result)
		fmt.Printf("Actual: %v\n\n", math.Sqrt(float64(num)))
		num += 1
	}
}
