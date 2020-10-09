package acceptance_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BooleanCat/alexandrium/types"
)

var _ = Describe("Acceptance", func() {
	var cmd *exec.Cmd

	BeforeEach(func() {
		cmd = exec.Command(binPath)
		cmd.Stdout = GinkgoWriter
		cmd.Stderr = GinkgoWriter
		Expect(cmd.Start()).To(Succeed())
		Eventually(ping(serverAddress)).Should(Succeed())
	})

	AfterEach(func() {
		Expect(cmd.Process.Signal(os.Interrupt)).To(Succeed())
		Expect(cmd.Wait()).To(Succeed())
	})

	Describe("GET /books/9781788547383", func() {
		It("responds with the correct book data by ISBN", func() {
			response := httpGet(serverAddress + "/books/9781788547383")
			defer closeIgnoreError(response.Body)

			Expect(response.StatusCode).To(Equal(http.StatusOK))

			var book types.Book
			Expect(json.NewDecoder(response.Body).Decode(&book)).To(Succeed())
			Expect(book).To(Equal(types.Book{
				ID:        "76341e07-911c-44fd-aafa-13b43daf3494",
				ISBN:      "9781788547383",
				Name:      "Cage of Souls",
				Publisher: "Head of Zeus",
				Authors:   []string{"ea1ff7d7-67cd-477c-8cb7-8756619e275d"},
			}))
		})
	})

	Describe("GET /books/76341e07-911c-44fd-aafa-13b43daf3494", func() {
		It("responds with the correct book data by Alexandrium ID", func() {
			response := httpGet(serverAddress + "/books/76341e07-911c-44fd-aafa-13b43daf3494")
			defer closeIgnoreError(response.Body)

			Expect(response.StatusCode).To(Equal(http.StatusOK))

			var book types.Book
			Expect(json.NewDecoder(response.Body).Decode(&book)).To(Succeed())
			Expect(book).To(Equal(types.Book{
				ID:        "76341e07-911c-44fd-aafa-13b43daf3494",
				ISBN:      "9781788547383",
				Name:      "Cage of Souls",
				Publisher: "Head of Zeus",
				Authors:   []string{"ea1ff7d7-67cd-477c-8cb7-8756619e275d"},
			}))
		})
	})

	Describe("GET /authors/ea1ff7d7-67cd-477c-8cb7-8756619e275d", func() {
		It("responds with the correct author data by Alexandrium ID", func() {
			response := httpGet(serverAddress + "/authors/ea1ff7d7-67cd-477c-8cb7-8756619e275d")
			defer closeIgnoreError(response.Body)

			var author types.Author
			Expect(json.NewDecoder(response.Body).Decode(&author)).To(Succeed())
			Expect(author).To(Equal(types.Author{
				ID:   "ea1ff7d7-67cd-477c-8cb7-8756619e275d",
				Name: "Adrian Tchaikovsky",
			}))
		})
	})
})

func ping(url string) func() error {
	return func() error {
		response, err := http.Get(url + "/ping")
		if err != nil {
			return err
		}
		if err = response.Body.Close(); err != nil {
			return err
		}

		if response.StatusCode == http.StatusNoContent {
			return nil
		}

		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
}

func closeIgnoreError(c io.Closer) {
	_ = c.Close()
}
