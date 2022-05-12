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

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
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

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// instantiate a custom http server.  do this to allow custom error logging to be used.
	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Http Server started and listening on http://%s ...", cfg.addr)
	// errorLog.Fatal(http.ListenAndServe(cfg.addr, mux))
	errorLog.Fatal(srv.ListenAndServe())

}
