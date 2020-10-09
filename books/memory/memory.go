package memory

import (
	"github.com/BooleanCat/alexandrium/books"
	"github.com/BooleanCat/alexandrium/types"
)

type Books struct{}

func (_ *Books) ByISBN(isbn string) (types.Book, error) {
	for _, book := range data {
		if book.ISBN == isbn {
			return book, nil
		}
	}

	return types.Book{}, books.NotFoundError{}
}

func (_ *Books) ByID(id string) (types.Book, error) {
	for _, book := range data {
		if book.ID == id {
			return book, nil
		}
	}

	return types.Book{}, books.NotFoundError{}
}

var _ books.Books = new(Books)

var data = []types.Book{
	{
		ID:        "76341e07-911c-44fd-aafa-13b43daf3494",
		ISBN:      "9781788547383",
		Name:      "Cage of Souls",
		Author:    "Adrian Tchaikovsky",
		Publisher: "Head of Zeus",
	},
}
