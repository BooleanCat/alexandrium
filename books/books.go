package books

import "errors"

type Books interface {
	ByISBN(string) (Book, error)
}

type Book struct {
	ISBN string `json:"isbn"`
	Name string `json:"name"`
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
