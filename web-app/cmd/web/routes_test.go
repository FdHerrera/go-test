package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_application_routes(t *testing.T) {
	registeredRoutes := []struct {
		registeredRoute string
		httpMethod      string
	}{
		{"/", "GET"},
		{"/static/*", "GET"},
	}

	var app application

	mux := app.routes()

	chiRoutes := mux.(chi.Routes)

	for _, e := range registeredRoutes {
		exists := routeExists(e.registeredRoute, e.httpMethod, chiRoutes)
		if !exists {
			t.Errorf("Route not registered: [%s]", e.registeredRoute)
		}
	}
}

func routeExists(testRoute string, testMethod string, chiRoutes chi.Routes) bool {
	found := false

	_ = chi.Walk(
		chiRoutes,
		func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
			if strings.EqualFold(method, testMethod) && strings.EqualFold(route, testRoute) {
				found = true
			}
			return nil
		},
	)
	return found
}
