package memory_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BooleanCat/alexandrium/store"
	"github.com/BooleanCat/alexandrium/store/memory"
)

var _ = Describe("Memory", func() {
	var authors *memory.AuthorStore

	BeforeEach(func() {
		authors = new(memory.AuthorStore)
	})

	Describe("ByID", func() {
		It("returns the correct author for an ID", func() {
			book, err := authors.ByID("ea1ff7d7-67cd-477c-8cb7-8756619e275d")
			Expect(err).NotTo(HaveOccurred())
			Expect(book.ID).To(Equal("ea1ff7d7-67cd-477c-8cb7-8756619e275d"))
		})

		When("an author isn't found", func() {
			It("returns a NotFound error", func() {
				_, err := authors.ByID("not-found")
				Expect(store.IsNotFound(err)).To(BeTrue())
			})
		})
	})
})
