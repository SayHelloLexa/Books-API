package main

import (
	"net/http"

	h "github.com/sayhellolexa/api-example/api/handlers"
	m "github.com/sayhellolexa/api-example/api/middleware"

	"github.com/gorilla/mux"
)

func main() {
	a := &h.Api{
		Router: mux.NewRouter(),
		Books:  &h.Books{},
	}

	a.Router.Use(m.JsonHeaderMiddleware)

	*a.Books = append(*a.Books, h.Book{
		ID:    "1",
		Title: "Властелин колец",
		Author: &h.Author{
			Firstname: "Джон",
			Lastname:  "Толкин",
		},
	})

	*a.Books = append(*a.Books, h.Book{
		ID:    "2",
		Title: "Преступление и наказание",
		Author: &h.Author{
			Firstname: "Федор",
			Lastname:  "Достоевский",
		},
	})

	a.Endpoints()
	http.ListenAndServe(":8080", a.Router)
}
