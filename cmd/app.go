package main

import (
	"log"
	"net/http"

	h "github.com/sayhellolexa/api-example/pkg/api"
)

func main() {
	a := h.New()

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

	log.Println("Сервер слушает :8080...")
	http.ListenAndServe(":8080", a.Router)
}
