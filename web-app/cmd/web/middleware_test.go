package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_addIpToContext(t *testing.T) {
	tests := []struct {
		headerName   string
		headerValue  string
		address      string
		emptyAddress bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.1.1.1", "", false},
		{"", "", "hello:world", false},
	}

	var app application

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, "not preset")
		}

		ip, ok := val.(string)
		if !ok {
			t.Error("not a string")
		}

		t.Log(ip)
	})

	for _, e := range tests {
		handlerToTest := app.addIpToContext(handler)

		req := httptest.NewRequest("GET", "http://testing", nil)

		if e.emptyAddress {
			req.RemoteAddr = ""
		}

		if len(e.headerName) > 0 {
			req.Header.Add(e.headerName, e.headerValue)
		}

		if len(e.address) > 0 {
			req.RemoteAddr = e.address
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}
