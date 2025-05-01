package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	m "github.com/sayhellolexa/api-example/pkg/middleware"
	"math/rand"
	"net/http"
	"strconv"
)

// Api - определение структуры API
type Api struct {
	Router *mux.Router
	Books  *Books
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

// New - конструктор API
func New() *Api {
	api := &Api{
		Router: mux.NewRouter(),
		Books:  &Books{},
	}
	api.endpoints()

	return api
}

// Endpoints - эндпоинты API
func (a *Api) endpoints() {
	a.Router.Use(m.JsonHeaderMiddleware)

	a.Router.HandleFunc("/books", a.Books.getBooks).Methods(http.MethodGet)
	a.Router.HandleFunc("/books/{id}", a.Books.getBook).Methods(http.MethodGet)
	a.Router.HandleFunc("/books", a.Books.createBook).Methods(http.MethodPost)
	a.Router.HandleFunc("/books/{id}", a.Books.updBook).Methods(http.MethodPut)
	a.Router.HandleFunc("/books/{id}", a.Books.deleteBook).Methods(http.MethodDelete)
}

// getBooks - получить все книги
func (b *Books) getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(b)
}

// getBook - получить книгу по ID
func (b *Books) getBook(w http.ResponseWriter, r *http.Request) {
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
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000))
	*b = append(*b, book)
	json.NewEncoder(w).Encode(book)
}

// updBook - обновить ID книги
func (b *Books) updBook(w http.ResponseWriter, r *http.Request) {
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
	params := mux.Vars(r)

	for i, item := range *b {
		if item.ID == params["id"] {
			*b = append((*b)[:i], (*b)[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(b)
}
