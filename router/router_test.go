package router_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BooleanCat/alexandrium/router"
	"github.com/BooleanCat/alexandrium/router/internal"
	"github.com/BooleanCat/alexandrium/store"
	"github.com/BooleanCat/alexandrium/types"
)

var _ = Describe("Router", func() {
	var (
		server      *httptest.Server
		fakeBooks   *internal.FakeBooks
		fakeAuthors *internal.FakeAuthors
	)

	BeforeEach(func() {
		fakeBooks = new(internal.FakeBooks)
		fakeAuthors = new(internal.FakeAuthors)
		server = httptest.NewServer(router.New(fakeBooks, fakeAuthors))
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

	Describe("/books/{id}", func() {
		var (
			response *http.Response
			id       string
		)

		BeforeEach(func() {
			id = "76341e07-911c-44fd-aafa-13b43daf3494"
			fakeBooks.ByIDReturns(types.Book{ID: id}, nil)
		})

		JustBeforeEach(func() {
			response = httpGet(server.URL + "/books/" + id)
		})

		AfterEach(func() {
			Expect(response.Body.Close()).To(Succeed())
		})

		It("responds", func() {
			By("having status code 200 OK", func() {
				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			By("having the correct Content-Type", func() {
				Expect(response.Header.Get("Content-Type")).To(Equal("application/json"))
			})

			By("having searched by ISBN", func() {
				Expect(fakeBooks.ByIDCallCount()).To(Equal(1))
			})

			By("delivering the book in the response body", func() {
				var book types.Book
				Expect(json.NewDecoder(response.Body).Decode(&book)).To(Succeed())
				Expect(book.ID).To(Equal(id))
			})
		})

		When("the book isn't found", func() {
			BeforeEach(func() {
				fakeBooks.ByIDReturns(types.Book{}, store.NotFoundError{})
			})

			It("responds", func() {
				By("having status code 404 Not Found", func() {
					Expect(response.StatusCode).To(Equal(http.StatusNotFound))
				})

				By("having the correct Content-Type", func() {
					Expect(response.Header.Get("Content-Type")).To(Equal("application/json"))
				})

				By("having searched by ID", func() {
					Expect(fakeBooks.ByIDCallCount()).To(Equal(1))
				})
			})
		})

		When("searching by an invalid ID", func() {
			BeforeEach(func() {
				id = "not-valid"
			})

			It("responds", func() {
				By("having status code 400 Bad Request", func() {
					Expect(response.StatusCode).To(Equal(http.StatusBadRequest))
				})

				By("having the correct Content-Type", func() {
					Expect(response.Header.Get("Content-Type")).To(Equal("application/json"))
				})

				By("not having searched by ID", func() {
					Expect(fakeBooks.ByIDCallCount()).To(BeZero())
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

				By("having the correct Content-Type", func() {
					Expect(response.Header.Get("Content-Type")).To(Equal("application/json"))
				})

				By("having searched by ID", func() {
					Expect(fakeBooks.ByIDCallCount()).To(Equal(1))
				})
			})
		})

		When("searching by ISBN", func() {
			BeforeEach(func() {
				id = "9781788547383"
				fakeBooks.ByISBNReturns(types.Book{ISBN: "9781788547383"}, nil)
			})

			It("responds", func() {
				By("having status code 200 OK", func() {
					Expect(response.StatusCode).To(Equal(http.StatusOK))
				})

				By("having the correct Content-Type", func() {
					Expect(response.Header.Get("Content-Type")).To(Equal("application/json"))
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
					fakeBooks.ByISBNReturns(types.Book{}, store.NotFoundError{})
				})

				It("responds", func() {
					By("having status code 404 Not Found", func() {
						Expect(response.StatusCode).To(Equal(http.StatusNotFound))
					})

					By("having the correct Content-Type", func() {
						Expect(response.Header.Get("Content-Type")).To(Equal("application/json"))
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

					By("having the correct Content-Type", func() {
						Expect(response.Header.Get("Content-Type")).To(Equal("application/json"))
					})

					By("having searched by ISBN", func() {
						Expect(fakeBooks.ByISBNCallCount()).To(Equal(1))
					})
				})
			})
		})
	})

	Describe("/authors/{id}", func() {
		var (
			response *http.Response
			id       string
		)

		BeforeEach(func() {
			id = "ea1ff7d7-67cd-477c-8cb7-8756619e275d"
			fakeAuthors.ByIDReturns(types.Author{ID: id}, nil)
		})

		JustBeforeEach(func() {
			response = httpGet(server.URL + "/authors/" + id)
		})

		AfterEach(func() {
			Expect(response.Body.Close()).To(Succeed())
		})

		It("responds", func() {
			By("having status code 200 OK", func() {
				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			By("having the correct Content-Type", func() {
				Expect(response.Header.Get("Content-Type")).To(Equal("application/json"))
			})

			By("having searched by ID", func() {
				Expect(fakeAuthors.ByIDCallCount()).To(Equal(1))
			})

			By("delivering the author in the response body", func() {
				var author types.Author
				Expect(json.NewDecoder(response.Body).Decode(&author)).To(Succeed())
				Expect(author.ID).To(Equal(id))
			})
		})

		When("the author isn't found", func() {
			BeforeEach(func() {
				fakeAuthors.ByIDReturns(types.Author{}, store.NotFoundError{})
			})

			It("responds", func() {
				By("having status code 404 Not Found", func() {
					Expect(response.StatusCode).To(Equal(http.StatusNotFound))
				})

				By("having the correct Content-Type", func() {
					Expect(response.Header.Get("Content-Type")).To(Equal("application/json"))
				})

				By("having searched by ID", func() {
					Expect(fakeAuthors.ByIDCallCount()).To(Equal(1))
				})
			})
		})

		When("searching by ID fails", func() {
			BeforeEach(func() {
				fakeAuthors.ByIDReturns(types.Author{}, errors.New("oops"))
			})

			It("responds", func() {
				By("having status code 500 Internal Server Error", func() {
					Expect(response.StatusCode).To(Equal(http.StatusInternalServerError))
				})

				By("having the correct Content-Type", func() {
					Expect(response.Header.Get("Content-Type")).To(Equal("application/json"))
				})

				By("having searched by ID", func() {
					Expect(fakeAuthors.ByIDCallCount()).To(Equal(1))
				})
			})
		})

		When("searching by an invalid ID", func() {
			BeforeEach(func() {
				id = "not-valid"
			})

			It("responds", func() {
				By("having status code 400 Bad Request", func() {
					Expect(response.StatusCode).To(Equal(http.StatusBadRequest))
				})

				By("having the correct Content-Type", func() {
					Expect(response.Header.Get("Content-Type")).To(Equal("application/json"))
				})

				By("having not having searched by ID", func() {
					Expect(fakeAuthors.ByIDCallCount()).To(BeZero())
				})
			})
		})
	})
})
