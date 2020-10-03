package router

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/BooleanCat/alexandrium/books"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate -o ./internal/fake_books.go github.com/BooleanCat/alexandrium/books.Books

// New creates alexandrium's HTTP router.
func New(b books.Books) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/ping", HandlePing)
	router.Get("/books/{isbn}", HandleBooksByISBN(b))

	return router
}

// HandlePing is a handler returning 204 No Content
func HandlePing(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// HandleBooksByISBN responds with a book by ISBN
func HandleBooksByISBN(b books.Books) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isbn := chi.URLParam(r, "isbn")

		book, err := b.ByISBN(isbn)
		if books.IsNotFound(err) {
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(&book); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

