package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.addIpToContext)
	mux.Use(app.Session.LoadAndSave)

	// register routes
	mux.Get("/", app.Home)
	mux.Post("/login", app.login)

	// static assets
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	return mux
}
