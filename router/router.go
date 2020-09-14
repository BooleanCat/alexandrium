package router

import (
	"net/http"

	"github.com/go-chi/chi"
)

// New creates alexandrium's HTTP router.
func New() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/ping", Ping)

	return router
}

// Ping is a handler returning 204 No Content
func Ping(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
