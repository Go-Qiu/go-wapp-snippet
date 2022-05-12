package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// define a new command-line flag, 'addr', default value of ":4000"
	addr := flag.String("addr", ":4000", "HTTP network address")

	// it is recommended not to use the default server mux in http package in production.
	// recommended to make a declaration and use for instantiating a http server.
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// fixed path patterns
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// subtree path pattern.
	// this pattern is a catch all, when incoming request uri does not
	// match any of the above fixed path patterns.
	mux.HandleFunc("/", home)

	log.Println("Http Server started and listening on http://localhost:4000 ...")
	log.Fatal(http.ListenAndServe(*addr, mux))

}
