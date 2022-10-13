package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Change the signature of the home handler
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		//http.NotFound(w, r)
		app.notFound(w)
		return
	}

	// Template files.
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	// template.ParseFiles() reads the template file into a template set.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		//log.Print(err.Error())

		/*app.errorLog.Print(err.Error())
		http.Error(w, "Internal Server Error", 500) // for error  */
		app.serverError(w, err)
		return
	}
	// Execute() method write the template content as the response body.
	// ExecuteTemplate() method writes the content of the "base".
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		//log.Print(err.Error())

		/*app.errorLog.Print(err.Error())
		http.Error(w, "Internal Server Error", 500) */
		app.serverError(w, err)
	}
}
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		//http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
