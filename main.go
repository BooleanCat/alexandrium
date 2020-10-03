package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BooleanCat/alexandrium/books"
	"github.com/BooleanCat/alexandrium/router"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router.New(booksImpl{}),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen and server: %v", err)
		}
	}()

	<-signals

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown: %v", err)
	}
}

type booksImpl struct {}

func (_ booksImpl) ByISBN(_ string) (books.Book, error) {
	return books.Book{}, books.NotFoundError{}
}

var _ books.Books = booksImpl{}