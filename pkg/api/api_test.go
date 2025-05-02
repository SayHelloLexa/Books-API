package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var api *Api

func TestMain(m *testing.M) {
	api = New()
	api.endpoints()
	os.Exit(m.Run())
}

func Test_createBook(t *testing.T) {
	data := Book{
		Title: "Энциклопедия Арнольда Шварцнеггера",
		Author: &Author{
			Firstname: "Арнольд",
			Lastname:  "Шварцнеггер",
		},
	}
	initLen := len(*api.Books)

	jsn, _ := json.Marshal(data)

	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(jsn))

	rr := httptest.NewRecorder()

	api.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен, получили: %v, хотели: %v", rr.Code, http.StatusOK)
	}
	t.Log("Response: ", rr.Body)

	if len(*api.Books) != initLen+1 {
		t.Errorf("книга не добавилась, получили: %d, хотели: %d", len(*api.Books), initLen+1)
	}
}

func Test_getBooks(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rr := httptest.NewRecorder()
	api.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен, получили: %v, хотели: %v", rr.Code, http.StatusOK)
	}
	t.Log("Response: ", rr.Body)
}

func Test_deleteBook(t *testing.T) {
	*api.Books = append(*api.Books, Book{
		ID:    "1",
		Title: "Книга для удаления",
		Author: &Author{
			Firstname: "Вася",
			Lastname:  "Залупкин",
		},
	})

	initLen := len(*api.Books)

	req := httptest.NewRequest(http.MethodDelete, "/books/1", nil)
	rr := httptest.NewRecorder()
	api.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен, получили: %v, хотели: %v", rr.Code, http.StatusOK)
	}
	t.Log("Response ", rr.Body)

	if len(*api.Books) != initLen-1 {
		t.Errorf("книга не удалилась, получили: %d, хотели: %d", len(*api.Books), initLen-1)
	}
}
