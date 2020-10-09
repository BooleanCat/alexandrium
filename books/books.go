package books

import (
	"errors"

	"github.com/BooleanCat/alexandrium/types"
)

type Books interface {
	ByISBN(string) (types.Book, error)
	ByID(string) (types.Book, error)
}

func IsNotFound(err error) bool {
	var e NotFoundError
	return errors.As(err, &e)
}

type NotFoundError struct{}

func (err NotFoundError) Error() string {
	return "not found"
}

var _ error = NotFoundError{}
