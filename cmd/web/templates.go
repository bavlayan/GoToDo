package main

import (
	"fmt"
	"path/filepath"
	"text/template"
	"time"

	"github.com/bavlayan/GoToDo/pkg/forms"
	"github.com/bavlayan/GoToDo/pkg/models"
)

type templateData struct {
	TodoItem          *models.TodoItems
	TodoItems         []*models.TodoItems
	Form              *forms.Form
	Flash             string
	AuthenticatedUser string
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println(name)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}
