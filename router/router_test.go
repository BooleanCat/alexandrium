package router_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BooleanCat/alexandrium/books"
	"github.com/BooleanCat/alexandrium/router"
	"github.com/BooleanCat/alexandrium/router/internal"
	"github.com/BooleanCat/alexandrium/types"
)

var _ = Describe("Router", func() {
	var (
		server    *httptest.Server
		fakeBooks *internal.FakeBooks
	)

	BeforeEach(func() {
		fakeBooks = new(internal.FakeBooks)
		server = httptest.NewServer(router.New(fakeBooks))
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("/ping", func() {
		It("responds with 204 No Content", func() {
			response := httpGet(server.URL + "/ping")
			Expect(response.Body.Close()).To(Succeed())
			Expect(response.StatusCode).To(Equal(http.StatusNoContent))
		})
	})

	Describe("/books/{isbn}", func() {
		var response *http.Response

		BeforeEach(func() {
			fakeBooks.ByISBNReturns(types.Book{ISBN: "9781788547383"}, nil)
		})

		JustBeforeEach(func() {
			response = httpGet(server.URL + "/books/9781788547383")
		})

		AfterEach(func() {
			Expect(response.Body.Close()).To(Succeed())
		})

		It("responds", func() {
			By("having status code 200 OK", func() {
				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			By("having searched by ISBN", func() {
				Expect(fakeBooks.ByISBNCallCount()).To(Equal(1))
			})

			By("delivering the book in the response body", func() {
				var book types.Book
				Expect(json.NewDecoder(response.Body).Decode(&book)).To(Succeed())
				Expect(book.ISBN).To(Equal("9781788547383"))
			})
		})

		When("the book isn't found", func() {
			BeforeEach(func() {
				fakeBooks.ByISBNReturns(types.Book{}, books.NotFoundError{})
			})

			It("responds", func() {
				By("having status code 404 Not Found", func() {
					Expect(response.StatusCode).To(Equal(http.StatusNotFound))
				})

				By("having searched by ISBN", func() {
					Expect(fakeBooks.ByISBNCallCount()).To(Equal(1))
				})
			})
		})

		When("searching by ISBN fails", func() {
			BeforeEach(func() {
				fakeBooks.ByISBNReturns(types.Book{}, errors.New("oops"))
			})

			It("responds", func() {
				By("having status code 500 Internal Server Error", func() {
					Expect(response.StatusCode).To(Equal(http.StatusInternalServerError))
				})

				By("having searched by ISBN", func() {
					Expect(fakeBooks.ByISBNCallCount()).To(Equal(1))
				})
			})
		})
	})

	Describe("/books/{id}", func() {
		var response *http.Response

		BeforeEach(func() {
			fakeBooks.ByIDReturns(types.Book{ID: "76341e07-911c-44fd-aafa-13b43daf3494"}, nil)
		})

		JustBeforeEach(func() {
			response = httpGet(server.URL + "/books/76341e07-911c-44fd-aafa-13b43daf3494")
		})

		AfterEach(func() {
			Expect(response.Body.Close()).To(Succeed())
		})

		It("responds", func() {
			By("having status code 200 OK", func() {
				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			By("having searched by ISBN", func() {
				Expect(fakeBooks.ByIDCallCount()).To(Equal(1))
			})

			By("delivering the book in the response body", func() {
				var book types.Book
				Expect(json.NewDecoder(response.Body).Decode(&book)).To(Succeed())
				Expect(book.ID).To(Equal("76341e07-911c-44fd-aafa-13b43daf3494"))
			})
		})

		When("the book isn't found", func() {
			BeforeEach(func() {
				fakeBooks.ByIDReturns(types.Book{}, books.NotFoundError{})
			})

			It("responds", func() {
				By("having status code 404 Not Found", func() {
					Expect(response.StatusCode).To(Equal(http.StatusNotFound))
				})

				By("having searched by ISBN", func() {
					Expect(fakeBooks.ByIDCallCount()).To(Equal(1))
				})
			})
		})

		When("searching by ID fails", func() {
			BeforeEach(func() {
				fakeBooks.ByIDReturns(types.Book{}, errors.New("oops"))
			})

			It("responds", func() {
				By("having status code 500 Internal Server Error", func() {
					Expect(response.StatusCode).To(Equal(http.StatusInternalServerError))
				})

				By("having searched by ISBN", func() {
					Expect(fakeBooks.ByIDCallCount()).To(Equal(1))
				})
			})
		})
	})
})
