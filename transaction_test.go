package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"testing"
)

func TestCreateEntry(t *testing.T) {

	var jsonStr = []byte(`{"Amount":365.78,"Currency":"CLP"}`)

	req, err := http.NewRequest("POST", "/api/PostTransaction", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `Done!`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}
