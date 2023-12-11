package main

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Has(t *testing.T) {
	form := NewForm(nil)

	has := form.Has("whatever")

	if has {
		t.Error("should not have")
	}

	postedData := url.Values{}

	postedData.Add("a", "a")
	form = NewForm(postedData)

	has = form.Has("a")

	if !has {
		t.Error("Should have registered urls")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/x", nil)

	form := NewForm(r.PostForm)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("Should not be valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	form = NewForm(postedData)
	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("Should be valid because form has the required values")
	}
}

func TestForm_Check(t *testing.T) {
	form := NewForm(nil)

	form.Check(false, "something", "something is required")

	if form.Valid() {
		t.Error("Should not be valid because Check got false")
	}
}

func TestForm_ErrorGet(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "something", "something is required")

	s := form.Errors.Get("something")

	if len(s) == 0 {
		t.Error("Should have got an error named 'something'")
	}

	s = form.Errors.Get("some error that should not be there")

	if s != "" {
		t.Error("Should not found an error not registered")
	}
}
