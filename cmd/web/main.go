package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type config struct {
	addr      string
	staticDir string
}

func main() {

	// define a new command-line flag, 'addr', default value of ":4000".
	// this is added to allow the server to be started on a specific network address and port
	// at the command-line (e.g. go run . -addr="192.168.1.3:8081")
	// addr := flag.String("addr", ":4000", "HTTP network address")
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address (and port)")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")

	// import to make this function call to make all of the above flags definition useable.
	flag.Parse()

	// declare custom loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

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
	// leave this out, to get '404', if there is no need for "/" handling, for security sack.
	mux.HandleFunc("/", home)

	// instantiate a custom http server.  do this to allow custom error logging to be used.
	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Http Server started and listening on http://%s ...", cfg.addr)
	// errorLog.Fatal(http.ListenAndServe(cfg.addr, mux))
	errorLog.Fatal(srv.ListenAndServe())

}
