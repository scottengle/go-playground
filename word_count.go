package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

// Determine word count for each word in the specified string
func WordCount(s string) map[string]int {

	// Create a new map with string keys and integer values
	word_map := make(map[string]int)

	// Iterate over the words in the original string
	for _, v := range strings.Fields(s) {

		// Increment the word count when a word is found
		word_map[v] += 1
	}

	return word_map
}

func main() {
	wc.Test(WordCount)
}
