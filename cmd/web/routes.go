package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// it is recommended not to use the default server mux in http package in production.
	// recommended to make a declaration and use for instantiating a http server.
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// fixed path patterns
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// subtree path pattern.
	// this pattern is a catch all, when incoming request uri does not
	// match any of the above fixed path patterns.
	// leave this out, to get '404', if there is no need for "/" handling, for security sack.
	mux.HandleFunc("/", app.home)
	return mux
}
