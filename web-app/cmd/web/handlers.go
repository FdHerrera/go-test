package main

import (
	"html/template"
	"net/http"
	"path"
)

var pathToTemplates string = "./templates/"

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "home.gohtml", &TemplateData{})
}

type TemplateData struct {
	Ip   string
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {
	// parse template
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t))
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return err
	}

	data.Ip = app.ipFromContext(r.Context())

	err = parsedTemplate.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}
