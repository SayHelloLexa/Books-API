package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Api - определение структуры API
type Api struct {
	router *mux.Router
	books  *Books
}

// Books - тип для хранения книг
type Books []Book

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Endpoints - эндпоинты API
func (a *Api) Endpoints() {
	a.router.HandleFunc("/books", a.books.getBooks).Methods(http.MethodGet)
	a.router.HandleFunc("/books/{id}", a.books.getBook).Methods(http.MethodGet)
	a.router.HandleFunc("/books", a.books.createBook).Methods(http.MethodPost)
	a.router.HandleFunc("/books/{id}", a.books.updBook).Methods(http.MethodPut)
	a.router.HandleFunc("/books/{id}", a.books.deleteBook).Methods(http.MethodDelete)
}

// getBooks - получить все книги
func (b *Books) getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

// getBook - получить книгу по ID
func (b *Books) getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, item := range *b {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// createBook - инициализировать книгу
func (b *Books) createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000))
	*b = append(*b, book)
	json.NewEncoder(w).Encode(book)
}

// updBook - обновить ID книги
func (b *Books) updBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range *b {
		if item.ID == params["id"] {
			*b = append((*b)[:i], (*b)[i+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			*b = append(*b, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(b)
}

// deleteBook - удалить книгу
func (b *Books) deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range *b {
		if item.ID == params["id"] {
			*b = append((*b)[:i], (*b)[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(b)
}
