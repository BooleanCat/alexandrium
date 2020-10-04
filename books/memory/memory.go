package memory

import "github.com/BooleanCat/alexandrium/books"

type Books struct{}

func (_ *Books) ByISBN(isbn string) (books.Book, error) {
	for _, book := range data {
		if book.ISBN == isbn {
			return book, nil
		}
	}

	return books.Book{}, books.NotFoundError{}
}

func (_ *Books) ByID(id string) (books.Book, error) {
	for _, book := range data {
		if book.ID == id {
			return book, nil
		}
	}

	return books.Book{}, books.NotFoundError{}
}

var _ books.Books = new(Books)

var data = []books.Book{
	{
		ID:        "76341e07-911c-44fd-aafa-13b43daf3494",
		ISBN:      "9781788547383",
		Name:      "Cage of Souls",
		Author:    "Adrian Tchaikovsky",
		Publisher: "Head of Zeus",
	},
}
