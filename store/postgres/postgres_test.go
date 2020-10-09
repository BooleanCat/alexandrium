package postgres_test

import (
	"context"

	"github.com/jackc/pgx/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BooleanCat/alexandrium/store/postgres"
	"github.com/BooleanCat/alexandrium/types"
)

var _ = Describe("Postgres", func() {
	var (
		connection *pgx.Conn
		authors    *postgres.AuthorStore
	)

	BeforeEach(func() {
		var err error
		connection, err = pgx.Connect(context.Background(), databaseURL)
		Expect(err).NotTo(HaveOccurred())

		authors = &postgres.AuthorStore{Connection: connection}
	})

	AfterEach(func() {
		Expect(connection.Close(context.Background())).To(Succeed())
	})

	It("does the thing", func() {
		author, err := authors.ByID("ea1ff7d7-67cd-477c-8cb7-8756619e275d")
		Expect(err).NotTo(HaveOccurred())
		Expect(author).To(Equal(types.Author{
			ID:   "ea1ff7d7-67cd-477c-8cb7-8756619e275d",
			Name: "Adrian Tchaikovsky",
		}))
	})
})
