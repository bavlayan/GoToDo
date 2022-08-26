package main

import (
	"fmt"
	"net/http"

	"github.com/bavlayan/GoToDo/pkg/forms"
	"github.com/bavlayan/GoToDo/pkg/models"
)

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display the user signup form..")
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new user")
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display the user login form...")
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate and login the user...")
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user...")
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	td_items, err := app.todoitems.GetDaily()
	if err != nil {
		app.serverError(w, err)
		return
	}

	success_message := app.session.PopString(r, "success_message")
	data := &templateData{
		TodoItems: td_items,
		Flash:     success_message,
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

	form := forms.New(r.PostForm)
	form.Require("itemdescription")

	item_description := r.PostForm.Get("itemdescription")

	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{
			Form: form,
		})
		return
	}

	_, err = app.todoitems.Save(item_description)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "success_message", "To Do Item successfully created!")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) createToDoItemFrom(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}
