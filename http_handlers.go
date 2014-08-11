package main

import (
	"fmt"
	"net/http"
)

type String string

type Struct struct {
	Greeting string

	Punct string

	Who string
}

func (str String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Send a structured message to the http response writer
	fmt.Fprint(w, str)
}

func (stc *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Send a structured message to the http response writer
	fmt.Fprintf(w, "%v%v %v", stc.Greeting, stc.Punct, stc.Who)
}

func main() {
	// Handle specific routes here
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})

	http.ListenAndServe("localhost:4000", nil)
}
