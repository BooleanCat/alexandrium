package router

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/BooleanCat/alexandrium/store"
	"github.com/BooleanCat/alexandrium/types"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate -o ./internal/fake_books.go github.com/BooleanCat/alexandrium/store.Books
//counterfeiter:generate -o ./internal/fake_authors.go github.com/BooleanCat/alexandrium/store.Authors

var isbn13Pattern = regexp.MustCompile(`\d{13}`)

// New creates alexandrium's HTTP router.
func New(books store.Books, authors store.Authors) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/ping", HandlePing)

	router.Get("/books/{id}", HandleGetBook(books))

	router.Get("/authors/{id}", HandleGetAuthor(authors))

	return router
}

// HandlePing is a handler returning 204 No Content
func HandlePing(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// HandleGetBook responds with a book by ISBN or ID
func HandleGetBook(books store.Books) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		w.Header().Set("Content-Type", "application/json")

		finder := getBookStrategy(books, id)
		if finder == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		book, err := finder(id)
		if store.IsNotFound(err) {
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

func HandleGetAuthor(authors store.Authors) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		w.Header().Set("Content-Type", "application/json")

		if _, err := uuid.Parse(id); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		author, err := authors.ByID(id)
		if store.IsNotFound(err) {
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(&author); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func getBookStrategy(books store.Books, id string) func(string) (types.Book, error) {
	if _, err := uuid.Parse(id); err == nil {
		return books.ByID
	}

	if isbn13Pattern.MatchString(id) {
		return books.ByISBN
	}

	return nil
}
