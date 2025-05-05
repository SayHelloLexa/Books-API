package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"math/rand"
	"net/http"
	"strconv"
)

// Api - определение структуры API
type Api struct {
	Store  *sessions.CookieStore
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
		Store:  sessions.NewCookieStore([]byte("32-byte-long-auth-key-123456789012")),
		Router: mux.NewRouter(),
		Books:  &Books{},
	}
	api.endpoints()

	return api
}

// Endpoints - эндпоинты API
func (api *Api) endpoints() {
	api.Router.Use(JsonHeaderMiddleware)
	//api.Router.Use(api.ApiSessionMiddleware)
	api.Router.Use(api.JwtMiddleware)

	api.Router.HandleFunc("/api/v1/authSession", api.authSession).Methods(http.MethodPost, http.MethodOptions)
	api.Router.HandleFunc("/api/v1/authJWT", api.authJWT).Methods(http.MethodPost, http.MethodOptions)

	api.Router.HandleFunc("/api/v1/books", api.Books.getBooks).Methods(http.MethodGet)
	api.Router.HandleFunc("/api/v1/books/{id}", api.Books.getBook).Methods(http.MethodGet)
	api.Router.HandleFunc("/api/v1/books", api.Books.createBook).Methods(http.MethodPost)
	api.Router.HandleFunc("/api/v1/books/{id}", api.Books.updBook).Methods(http.MethodPut)
	api.Router.HandleFunc("/api/v1/books/{id}", api.Books.deleteBook).Methods(http.MethodDelete)
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
