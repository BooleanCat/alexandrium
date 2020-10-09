package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/BooleanCat/alexandrium/store"
	"github.com/BooleanCat/alexandrium/types"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate -o ./internal/fake_books.go github.com/BooleanCat/alexandrium/books.Books

// New creates alexandrium's HTTP router.
func New(books store.Books) *chi.Mux {
	router := chi.NewRouter()

	uuid4Pattern := "[a-f0-9]{8}-?[a-f0-9]{4}-?4[a-f0-9]{3}-?[89ab][a-f0-9]{3}-?[a-f0-9]{12}"

	router.Get("/ping", HandlePing)

	router.Route("/books", func(r chi.Router) {
		r.Get(fmt.Sprintf("/{id:%s}", uuid4Pattern), HandleGetBook(books))
		r.Get("/{isbn}", HandleGetBookByISBN(books))
	})

	return router
}

// HandlePing is a handler returning 204 No Content
func HandlePing(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// HandleGetBook responds with a book by ISBN
func HandleGetBook(books store.Books) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		handleFindBook(books.ByID, id, w)
	}
}

// HandleGetBookByISBN responds with a book by ISBN
func HandleGetBookByISBN(books store.Books) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isbn := chi.URLParam(r, "isbn")
		handleFindBook(books.ByISBN, isbn, w)
	}
}

func handleFindBook(finder func(string) (types.Book, error), key string, w http.ResponseWriter) {
	book, err := finder(key)
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
