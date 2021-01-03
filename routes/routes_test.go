package routes

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMarkdown(t *testing.T){
	entry:="abc"
	result:= parseMarkdown([]byte(entry))
	if result == nil {
		t.Errorf("parseMarkdown should generate an output")
	}
}

// https://blog.questionable.services/article/testing-http-handlers-go/
func TestFindHome(t *testing.T){
	req, err := http.NewRequest("GET", "/wiki/home", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleGetWiki)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestFindUnknown(t *testing.T){
	req, err := http.NewRequest("GET", "/wiki/abc", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleGetWiki)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestCreate(t *testing.T){
	body:=bytes.NewBuffer([]byte("abc"))
	req, err := http.NewRequest("POST", "/wiki/testCreate", body)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandlePostWiki)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
