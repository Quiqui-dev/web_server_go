package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handleHello)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code. Got %v, wanted %v", status, http.StatusOK)
	}

	expectedBody := "<h1>Hello World</h1><div><p>testingggg</p></div>"

	if rr.Body.String() != expectedBody {
		t.Errorf("handler returnde unexpected body: got %v, wanted %v", rr.Body.String(), expectedBody)
	}
}
