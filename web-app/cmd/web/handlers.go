package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

var pathToTemplates string = "./templates/"

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var td = make(map[string]any)

	if app.Session.Exists(r.Context(), "test") {
		msg := app.Session.GetString(r.Context(), "test")
		td["test"] = msg
	} else {
		app.Session.Put(r.Context(), "test", "Hit this page at " + time.Now().UTC().String())
	}
	_ = app.render(w, r, "home.gohtml", &TemplateData{Data: td})
}

type TemplateData struct {
	Ip   string
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {
	// parse template
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t), path.Join(pathToTemplates, "baselayout.gohtml"))
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

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("err")
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	form := NewForm(r.PostForm)
	form.Required("email", "password")

	if !form.Valid() {
		fmt.Fprint(w, "failed validation")
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	log.Println(email, password)
	
	fmt.Fprint(w, email)
}
