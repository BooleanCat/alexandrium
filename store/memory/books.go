package memory

import (
	"github.com/BooleanCat/alexandrium/store"
	"github.com/BooleanCat/alexandrium/types"
)

type BookStore struct{}

func (_ *BookStore) ByISBN(isbn string) (types.Book, error) {
	for _, book := range bookData {
		if book.ISBN == isbn {
			return book, nil
		}
	}

	return types.Book{}, store.NotFoundError{}
}

func (_ *BookStore) ByID(id string) (types.Book, error) {
	for _, book := range bookData {
		if book.ID == id {
			return book, nil
		}
	}

	return types.Book{}, store.NotFoundError{}
}

var _ store.Books = new(BookStore)

var bookData = []types.Book{
	{
		ID:        "76341e07-911c-44fd-aafa-13b43daf3494",
		ISBN:      "9781788547383",
		Name:      "Cage of Souls",
		Authors:   []string{"ea1ff7d7-67cd-477c-8cb7-8756619e275d"},
		Publisher: "Head of Zeus",
	},
}
