package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// serverError() helper writes an error message and stack trace to the errorLog,
// send  500 Internal Server Error
func (app *application) serverError(w http.ResponseWriter, err error) {
	// debug.Stack - to get a stack trace
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Print(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError() helper sends a specific status code and corresponding description to the user.
func (app *application) clientError(w http.ResponseWriter, status int) {
	// http.StatusText() - to automatically
	//generate a human-friendly text representation of a given HTTP status code.
	// http.StatusText(400) == "Bad Request"
	http.Error(w, http.StatusText(status), status)
}

//	also notFound() helper - convenience wrapper around clientError
//
// (404 Not Found)
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
