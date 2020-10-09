package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/BooleanCat/alexandrium/books"
	"github.com/BooleanCat/alexandrium/types"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate -o ./internal/fake_books.go github.com/BooleanCat/alexandrium/books.Books

// New creates alexandrium's HTTP router.
func New(b books.Books) *chi.Mux {
	router := chi.NewRouter()

	uuid4Pattern := "[a-f0-9]{8}-?[a-f0-9]{4}-?4[a-f0-9]{3}-?[89ab][a-f0-9]{3}-?[a-f0-9]{12}"

	router.Get("/ping", HandlePing)
	router.Get(fmt.Sprintf("/books/{id:%s}", uuid4Pattern), HandleGetBook(b))
	router.Get("/books/{isbn}", HandleGetBookByISBN(b))

	return router
}

// HandlePing is a handler returning 204 No Content
func HandlePing(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// HandleGetBook responds with a book by ISBN
func HandleGetBook(b books.Books) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		handleFindBook(b.ByID, id, w)
	}
}

// HandleGetBookByISBN responds with a book by ISBN
func HandleGetBookByISBN(b books.Books) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isbn := chi.URLParam(r, "isbn")
		handleFindBook(b.ByISBN, isbn, w)
	}
}

func handleFindBook(finder func(string) (types.Book, error), key string, w http.ResponseWriter) {
	book, err := finder(key)
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
