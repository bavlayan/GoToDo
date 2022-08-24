package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	goToDo_middleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamic_middleware := alice.New(app.session.Enable)

	mux := pat.New()
	mux.Get("/", dynamic_middleware.ThenFunc(http.HandlerFunc(app.home)))
	mux.Get("/todoitems/item-details/:id", dynamic_middleware.ThenFunc(http.HandlerFunc(app.showToDoItemDetails)))
	mux.Post("/todoitems/create-todo-item", dynamic_middleware.ThenFunc(http.HandlerFunc(app.createToDoItem)))
	mux.Get("/todoitems/create-todo-item", dynamic_middleware.ThenFunc(http.HandlerFunc(app.createToDoItemFrom)))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return goToDo_middleware.Then(mux)
}
