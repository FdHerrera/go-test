package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	var tests = []struct {
		name               string
		endpoint           string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"404", "/fish", http.StatusNotFound},
	}

	routes := app.routes()

	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	pathToTemplates = "./../../templates/"

	for _, e := range tests {
		resp, err := ts.Client().Get(ts.URL + e.endpoint)
		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("For: [%s] - Expected [%d], bug got: [%d]", e.name, e.expectedStatusCode, resp.StatusCode)
		}
	}
}

func TestApp_Home(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req = addContextAndSession(req, app)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(app.Home)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Home should return status ok, got: [%d]", rr.Code)
	}

	body, _ := io.ReadAll(rr.Body)
	if !strings.Contains(string(body), "<small>From Session:") {
		t.Error("Got a wrong home")
	}
}

func addContextAndSession(req *http.Request, app application) *http.Request {
	ctx := context.WithValue(req.Context(), contextUserKey, "something")
	req = req.WithContext(ctx)

	ctx, _ = app.Session.Load(req.Context(), req.Header.Get("X-Session"))

	return req.WithContext(ctx)
}
