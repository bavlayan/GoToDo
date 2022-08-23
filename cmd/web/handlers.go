package main

import (
	"net/http"
	"strings"

	"github.com/bavlayan/GoToDo/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	td_items, err := app.todoitems.GetDaily()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{
		TodoItems: td_items,
	}

	app.render(w, r, "home.page.tmpl", data)
}

func (app *application) showToDoItemDetails(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get(":id")
	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	td, err := app.todoitems.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{
		TodoItem: td,
	}

	app.render(w, r, "show.page.tmpl", data)
}

func (app *application) createToDoItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	errors := make(map[string]string)

	item_description := r.PostForm.Get("itemdescription")

	if strings.TrimSpace(item_description) == "" {
		errors["itemdescription"] = "This field cannot be blank"
	}

	if len(errors) > 0 {
		app.render(w, r, "create.page.tmpl", &templateData{
			FormErrors: errors,
			FormData:   r.PostForm,
		})
		return
	}

	_, err = app.todoitems.Save(item_description)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) createToDoItemFrom(w http.ResponseWriter, r *http.Request) {
	errors := make(map[string]string)
	app.render(w, r, "create.page.tmpl", &templateData{
		FormErrors: errors,
	})
}
