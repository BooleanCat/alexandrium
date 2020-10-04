package books

import "errors"

type Books interface {
	ByISBN(string) (Book, error)
	ByID(string) (Book, error)
}

type Book struct {
	ID        string `json:"id"`
	ISBN      string `json:"isbn"`
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
	Author    string `json:"author"`
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
