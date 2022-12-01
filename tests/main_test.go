package main

import (
	"librarySystem/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAvailableBooks(t *testing.T) {
	req, err := http.NewRequest("GET", "/library/availableBooks", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("userid", "2df74fa7-2a32-4be5-871d-db677bcd7ea7")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(services.BookService{}.GetAvailableBooks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// func TestGetBooksByNameOrAuthorOrGenre(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/library", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	q := req.URL.Query()
// 	params := mux.Vars(request)
// 	userid := params["userid"]
// 	bookname := request.URL.Query().Get("bookname")
// 	authorname := request.URL.Query().Get("authorname")
// 	genre := request.URL.Query().Get("genre")
// 	q.Add("userid", "82fef933-ced2-4ab3-8445-978f407e6eed")
// 	q.Add("bookname","")
// 	req.URL.RawQuery = q.Encode()
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(services.BookService{}.GetBooksByNameOrAuthorOrGenre)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
// }
