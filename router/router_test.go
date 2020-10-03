package router_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BooleanCat/alexandrium/books"
	"github.com/BooleanCat/alexandrium/router"
	"github.com/BooleanCat/alexandrium/router/internal"
)

var _ = Describe("Router", func() {
	var (
		server *httptest.Server
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
			Expect(response.StatusCode).To(Equal(http.StatusNoContent))
		})
	})

	Describe("/books/{isbn}", func() {
		var response *http.Response

		BeforeEach(func() {
			fakeBooks.ByISBNReturns(books.Book{ISBN: "9781788547383"}, nil)
		})

		JustBeforeEach(func() {
			response = httpGet(server.URL + "/books/9781788547383")
		})

		AfterEach(func() {
			Expect(response.Body.Close()).To(Succeed())
		})

		It("succeeds", func() {
			By("responding with 200 OK", func() {
				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			By("responding with the requested book", func() {
				var book books.Book
				Expect(json.NewDecoder(response.Body).Decode(&book)).To(Succeed())
				Expect(book.ISBN).To(Equal("9781788547383"))
			})
		})

		When("the book isn't found", func() {
			BeforeEach(func() {
				fakeBooks.ByISBNReturns(books.Book{}, books.NotFoundError{})
			})

			It("responds with 404 Not Found", func() {
				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})
})
