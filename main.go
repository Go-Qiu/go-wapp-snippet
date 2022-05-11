package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.RequestURI != "/" {
		// exception handling.  return 404 when request uri does not match
		// any of the fixed path patterns.
		http.NotFound(w, r)
		return
	}
	// incoming request uri matches the pattern.
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Display a specific snippet ...</h1>"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	// 'POST' request check.
	if r.Method != "POST" {
		// indicate in the response header that the allow method is 'POST'
		w.Header().Set("Allow", "POST")

		// set the response code (i.e. 405) in the header
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

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
