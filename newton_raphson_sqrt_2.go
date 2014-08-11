package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(f float64) (float64, int, error) {

	if f < 0 {

		// Throw an error if f is negative
		return 0, 0, ErrNegativeSqrt(f)

	} else if f == 0 {

		// Special case for x = 0
		return 0, 1, nil

	} else {

		// Keep track of iterations
		i := 0

		// Start with 1 (no dividing by zero!)
		z := float64(1)

		// Internal variable to calculate deltas between iterations
		last := z

		for {
			z = z - (((z * z) - f) / (2 * z))

			// Increment the iteration counter
			i++

			// Exit out when the delta of change is sufficiently low
			if math.Abs(z-last) < 10e-16 {
				break
			}

			last = z
		}

		return z, i, nil

	}
}

// A small utility function to help with output formatting.
func resultHandler(num float64) {
	result, iter, err := Sqrt(num)

	if err == nil {
		fmt.Printf("Sqrt(%v), Approximated in %v Iterations\n", num, iter)
		fmt.Printf("Approx: %v\n", result)
		fmt.Printf("Actual: %v\n\n", math.Sqrt(float64(num)))
	} else {
		fmt.Printf("Sqrt(%v), threw an error:\n", num)
		fmt.Printf("%v\n\n", err)
	}
}

func main() {
	fmt.Println("Approximate the square root of a number using the Newton–Raphson method with error handling.\n")

	for i := -2; i < 3; i++ {
		resultHandler(float64(i))
	}
}
