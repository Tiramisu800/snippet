package main

import (
	"github.com/julienschmidt/httprouter" // New import
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	// Initialize the router.
	router := httprouter.New()

	//Wrapping notFound() helper for matching error message [ view/99 | /missing]
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w) //if router != "/"
	})

	//static files route
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	//middleware chain containing the middleware specific to our dynamic application routes
	dynamic := alice.New(app.sessionManager.LoadAndSave)

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.snippetView))
	router.Handler(http.MethodGet, "/snippet/create", dynamic.ThenFunc(app.snippetCreate))
	router.Handler(http.MethodPost, "/snippet/create", dynamic.ThenFunc(app.snippetCreatePost))

	//  middleware chain
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(router)
}
