package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Hold the application-wide dependencies // new application
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	// 2 loggers
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// init a new application struct
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Init a new http.Server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// Using new 2 loggers
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe() // call new http.Server struct
	errorLog.Fatal(err)
}
