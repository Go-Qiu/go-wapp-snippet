package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.RequestURI != "/" {
		// exception handling.  return 404 when request uri does not match
		// any of the fixed path patterns.  recommended approach for security sack.
		// http.NotFound(w, r)
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		// app.errorLog.Println(err.Error())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		// log.Println(err.Error())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, err)
		return
	}

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	// get the id parameter from the url query string
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	// 404 handling
	if err != nil || id < 1 {
		// http.NotFound(w, r)
		app.notFound(w)
		return
	}

	// ok.

	// w.Write([]byte("<h1>Display a specific snippet ...</h1>"))
	fmt.Fprintf(w, "Display a specific snippet with ID %d ...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	// 'POST' request check.
	if r.Method != http.MethodPost {
		// indicate in the response header that the allow method is 'POST'
		w.Header().Set("Allow", http.MethodPost)

		// set the response code (i.e. 405) in the header
		w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		// http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("<h1>Create a new snippet ...</h1>"))

}
