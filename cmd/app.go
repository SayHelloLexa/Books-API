package main

import (
	"log"
	"net/http"

	"github.com/sayhellolexa/api-example/pkg/api"
)

func main() {
	a := api.New()

	*a.Books = append(*a.Books, api.Book{
		ID:    "1",
		Title: "Властелин колец",
		Author: &api.Author{
			Firstname: "Джон",
			Lastname:  "Толкин",
		},
	})

	*a.Books = append(*a.Books, api.Book{
		ID:    "2",
		Title: "Преступление и наказание",
		Author: &api.Author{
			Firstname: "Федор",
			Lastname:  "Достоевский",
		},
	})

	log.Println("Сервер слушает :8080...")
	http.ListenAndServe(":8080", a.Router)
}
