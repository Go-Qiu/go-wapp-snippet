package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.RequestURI == "/" {
		// incoming request uri matches the pattern.
		w.Write([]byte("Hello from Snippetbox"))
	} else {
		// incoming request uri does not match the pattern.
		// request arrives here because of the catch all pattern behaviour.

		msg := fmt.Sprintln("<h1>BAD REQUEST</h1>")
		msg += fmt.Sprintf("<h3>Cannot find URI, '%s'</h3>", r.RequestURI)
		fmt.Fprintln(w, msg)
	}

}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Display a specific snippet ...</h1>"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Create a new snippet ...</h1>"))

}

func main() {

	// it is recommended not to use the default server mux in http package in production.
	// recommended to make a declaration and use for instantiating a http server.
	mux := http.NewServeMux()

	// fixed path patterns
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// subtree path pattern.
	// this pattern is a catch all, when incoming request uri does not
	// match any of the above fixed path patterns.
	mux.HandleFunc("/", home)

	log.Println("Http Server started and listening on http://localhost:4000 ...")
	log.Fatal(http.ListenAndServe(":4000", mux))

}
