package memory

import "github.com/BooleanCat/alexandrium/books"

type Books struct{}

func (_ *Books) ByISBN(isbn string) (books.Book, error) {
	if book, ok := data[isbn]; ok {
		return book, nil
	}

	return books.Book{}, books.NotFoundError{}
}

var _ books.Books = new(Books)

var data = map[string]books.Book{
	"9781788547383": {ISBN: "9781788547383", Name: "Cage of Souls"},
}