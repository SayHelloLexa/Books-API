package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	a := &Api{
		router: mux.NewRouter(),
		books:  &Books{},
	}

	*a.books = append(*a.books, Book{
		ID:    "1",
		Title: "Властелин колец",
		Author: &Author{
			Firstname: "Джон",
			Lastname:  "Толкин",
		},
	})

	*a.books = append(*a.books, Book{
		ID:    "2",
		Title: "Преступление и наказание",
		Author: &Author{
			Firstname: "Федор",
			Lastname:  "Достоевский",
		},
	})

	a.Endpoints()
	http.ListenAndServe(":8080", a.router)
}
