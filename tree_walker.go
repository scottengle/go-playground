package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

// Walker walks the tree t sending all values
// from the tree to the channel ch.
func Walker(t *tree.Tree, ch chan int) {

	if t == nil {
		return
	}

	Walker(t.Left, ch)
	ch <- t.Value
	Walker(t.Right, ch)
}

// The TreeWalker orchestrates the walking of the tree
// and returns a read-only channel for comparison later.
func TreeWalker(t *tree.Tree) <-chan int {

	// We make the channel writable at first
	ch := make(chan int)

	// Kick off the walker in its own goroutine.
	// Close the channel once we're done walking it.
	go func() {
		Walker(t, ch)
		close(ch)
	}()

	// This casts the channel to a read-only channel
	return ch
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	// Create two TreeWalker channels to receive values on
	c1, c2 := TreeWalker(t1), TreeWalker(t2)

	for {
		// This is a two value form of recieving from a channel.
		// The _ok values are false when all values have been received
		// from the channel and the channel is closed.
		left, left_ok := <-c1
		right, right_ok := <-c2

		if !left_ok || !right_ok {
			// If one of the channels have been closed, compare
			return left_ok == right_ok
		}

		if left != right {
			// Just exit out of the for loop if either of the incoming
			// values don't match.
			break
		}
	}

	return false
}

func main() {

	fmt.Printf("tree.New(1) same as tree.New(1)? %v\n\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("tree.New(1) same as tree.New(2)? %v\n\n", Same(tree.New(1), tree.New(2)))
}
