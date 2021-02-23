package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpHandler(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder, which satisfies the http.ResponseWriter interface, to record the response
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	// Our handler satisfies http.Handler, so we can call the ServeHTTP method directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(r, req)

	// Check that the status code is what we expect.
	if r.Code != http.StatusOK {
		t.Errorf("helloHandler returned status code %v, want %v", r.Code, http.StatusOK)
	}

	// Check that the response body is what we expect.
	want := "Hello, friend :) \n"
	got := r.Body.String()
	if got != want {
		t.Errorf("helloHandler returned body %q want %q", got, want)
	}
}
