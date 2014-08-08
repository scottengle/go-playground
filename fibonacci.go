package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {

	// Formula: f(n) = f(n-1) + f(n-2)
	//  where f(0) = 0 and f(1) = 1

	// Because we don't have function parameters,
	// we need to track n for f(n)

	n, n_minus_1, n_minus_2 := 0, 0, 0

	// Return a closure bound to n, n_minus_1 and n_minus_2
	return func() int {

		result := 0

		if n == 0 {
			// Special case for f(0) = 0
			result = 0
		} else if n == 1 {
			// Special case for f(1) = 1
			result = 1
		} else {
			// Calculate f(n)
			result = n_minus_1 + n_minus_2
		}

		// Remap the sequence values
		// result >> n_minus_1 >> n_minus_2
		n_minus_2 = n_minus_1
		n_minus_1 = result

		// Lastly, increment the current counter
		n += 1

		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("f(%v) = %v\n", i, f())
	}
}
