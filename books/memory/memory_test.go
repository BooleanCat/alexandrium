package memory_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BooleanCat/alexandrium/books"
	"github.com/BooleanCat/alexandrium/books/memory"
)

var _ = Describe("Memory", func() {
	var memoryBooks *memory.Books

	BeforeEach(func() {
		memoryBooks = new(memory.Books)
	})

	Describe("ByISBN", func() {
		It("returns the correct book for an ISBN", func() {
			book, err := memoryBooks.ByISBN("9781788547383")
			Expect(err).NotTo(HaveOccurred())
			Expect(book.ISBN).To(Equal("9781788547383"))
		})

		When("a book isn't found", func() {
			It("returns a NotFound error", func() {
				_, err := memoryBooks.ByISBN("not-found")
				Expect(books.IsNotFound(err)).To(BeTrue())
			})
		})
	})

	Describe("ByID", func() {
		It("returns the correct book for an ID", func() {
			book, err := memoryBooks.ByID("76341e07-911c-44fd-aafa-13b43daf3494")
			Expect(err).NotTo(HaveOccurred())
			Expect(book.ID).To(Equal("76341e07-911c-44fd-aafa-13b43daf3494"))
		})

		When("a book isn't found", func() {
			It("returns a NotFound error", func() {
				_, err := memoryBooks.ByID("not-found")
				Expect(books.IsNotFound(err)).To(BeTrue())
			})
		})
	})
})
