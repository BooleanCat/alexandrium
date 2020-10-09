package postgres_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/jackc/pgx/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var databaseURL string

func TestPostgres(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Postgres Suite")
}

var _ = BeforeSuite(func() {
	databaseURL = mustGetEnv("DATABASE_URL")

	connection, err := pgx.Connect(context.Background(), databaseURL)
	Expect(err).NotTo(HaveOccurred())
	defer func() { _ = connection.Close(context.Background()) }()

	query, err := ioutil.ReadFile("schema.sql")
	Expect(err).NotTo(HaveOccurred())

	_, err = connection.Exec(context.Background(), string(query))
	Expect(err).NotTo(HaveOccurred())
})

func mustGetEnv(key string) string {
	env := os.Getenv(key)
	if env == "" {
		Fail(fmt.Sprintf(`env "%s" must be set`, key))
	}

	return env
}
