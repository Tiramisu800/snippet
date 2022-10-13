package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Static issues
	// http.Dir function is relative to the project directory
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// mux.Handle() function register the file server as the handler
	// "/static" - prefix
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// swap the route declaration with new app
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
