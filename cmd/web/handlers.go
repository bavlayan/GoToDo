package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/bavlayan/GoToDo/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	td_items, err := app.todoitems.GetDaily()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, td_item := range td_items {
		fmt.Fprintf(w, "%v\n", td_item)
	}

	/*files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}*/
}

func (app *application) showToDoItemDetails(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
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

	data := &templateData{TodoItem: td}

	files := []string{
		"./ui/html/show.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	tf, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = tf.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) createToDoItemDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	description := "test from GoToDo app"
	affected_row, err := app.todoitems.Save(description)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/item-details?id=%d", affected_row), http.StatusSeeOther)

	w.Write([]byte("Create ToDO item"))
}
